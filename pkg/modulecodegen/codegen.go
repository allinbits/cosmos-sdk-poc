package modulecodegen

import (
	"fmt"
	"strings"

	"github.com/fdymylja/tmos/core/modulegen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

const metaImportPackage = protogen.GoImportPath("github.com/fdymylja/tmos/runtime/meta")
const moduleImportPackage = protogen.GoImportPath("github.com/fdymylja/tmos/runtime/module")

const ClientFileSuffix = ".clientset.starport.go"
const ObjectsFileSuffix = ".starport.go"

func PluginRunner(plugin *protogen.Plugin) error {
	groups := make(map[protogen.GoImportPath][]*protogen.File)
	for _, f := range plugin.Files {
		groups[f.GoImportPath] = append(groups[f.GoImportPath], f)
	}
	// then we parse single groups
	for _, group := range groups {
		err := genFile(plugin, group)
		if err != nil {
			return err
		}
	}
	return nil
}

func genFile(gen *protogen.Plugin, group []*protogen.File) error {
	for _, file := range group {
		genObjects(file, gen)
	}
	return nil
}

func genObjects(file *protogen.File, gen *protogen.Plugin) {

	if !meetsRequirements(file) {
		return
	}

	filename := fmt.Sprintf("%s%s", file.GeneratedFilenamePrefix, ObjectsFileSuffix)
	objectsFile := gen.NewGeneratedFile(filename, file.GoImportPath)
	objectsFile.P("package ", file.GoPackageName)

	var stateTransitions []*protogen.Message
	var stateObjects []*protogen.Message
	for _, msg := range file.Messages {
		md := msg.Desc
		// check if message option is present
		messageOptions := md.Options().(*descriptorpb.MessageOptions)
		isStateObject := proto.GetExtension(messageOptions, modulegen.E_StateObject).(bool)
		processed := false
		if isStateObject {
			genStateObject(objectsFile, msg)
			processed = true
			stateObjects = append(stateObjects, msg)
		}
		isStateTransition := proto.GetExtension(messageOptions, modulegen.E_StateTransition).(bool)
		if isStateTransition {
			if processed {
				gen.Error(fmt.Errorf("%s is defined as state object and state transition too which is not allowed", msg.Desc.Name()))
			}
			genStateTransition(objectsFile, msg)
			stateTransitions = append(stateTransitions, msg)
		}
	}
	// gen clientset
	genClientSet(objectsFile, stateObjects, stateTransitions)
}

func genClientSet(g *protogen.GeneratedFile, objects []*protogen.Message, transitions []*protogen.Message) {

	// gen constructor
	g.P("func NewClientSet(client ", moduleImportPackage.Ident("Client"), ") ClientSet {")
	g.P("return clientSet{client: client}")
	g.P("}")
	// gen clientset interface
	g.P("type ClientSet interface {")
	// add state objects client interface
	for _, obj := range objects {
		switch strings.HasSuffix(obj.GoIdent.GoName, "s") {
		case false:
			g.P(obj.GoIdent, "s()", " ", obj.GoIdent, "Client")
		case true:
			g.P(obj.GoIdent, "()", " ", obj.GoIdent, "Client")
		}
	}
	// gen state transitions interface
	for _, t := range transitions {
		g.P("Exec", t.GoIdent.GoName, "(msg *", t.GoIdent.GoName, ") error")
	}
	g.P("}")
	g.P()
	// gen client set concrete type
	g.P("type clientSet struct {")
	g.P("client ", moduleImportPackage.Ident("Client"))
	g.P("}")
}

func meetsRequirements(file *protogen.File) bool {
	for _, msg := range file.Messages {
		md := msg.Desc
		// check if message option is present
		messageOptions := md.Options().(*descriptorpb.MessageOptions)
		isStateObject := proto.GetExtension(messageOptions, modulegen.E_StateObject).(bool)
		if isStateObject {
			return true
		}
		isStateTransition := proto.GetExtension(messageOptions, modulegen.E_StateTransition).(bool)
		if isStateTransition {
			return true
		}
	}
	return false
}

func genStateTransition(g *protogen.GeneratedFile, message *protogen.Message) {
	// add state transition interface
	g.Import(metaImportPackage)
	g.P("func (x *", message.GoIdent, ") StateTransition() {}")
	g.P()
	g.P("func (x *", message.GoIdent, ") New() ", metaImportPackage.Ident("StateTransition"), " {")
	g.P("return new(", message.GoIdent, ")")
	g.P("}")
	g.P()
}

func genStateObject(g *protogen.GeneratedFile, message *protogen.Message) {
	g.Import(metaImportPackage)
	g.P("func (x *", message.GoIdent, ") StateObject() {}")
	g.P()
	g.P("func (x *", message.GoIdent, ") New() ", metaImportPackage.Ident("StateObject"), " {")
	g.P("return new(", message.GoIdent, ")")
	g.P("}")
	g.P()

	genClient(g, message)
}

func genClient(g *protogen.GeneratedFile, message *protogen.Message) {
	singleTon, primaryKey, primaryKeyGoType, err := parseSaveInfo(message)
	if err != nil {
		panic(err) // TODO not with panic
	}
	g.P("type ", message.GoIdent, "Client interface {")
	switch singleTon {
	case true:
		g.P("Get() (*", message.GoIdent, ", error)")
	case false:
		g.P("Get(", primaryKey, " ", primaryKeyGoType, ") (*", message.GoIdent, ", error)")
	}
	g.P("Create(", toLowerCamelCase(message.GoIdent), " *", message.GoIdent, ") error")
	g.P("Delete(", toLowerCamelCase(message.GoIdent), " *", message.GoIdent, ") error")
	g.P("Update(", toLowerCamelCase(message.GoIdent), " *", message.GoIdent, ") error")
	g.P("}")
	g.P()
	// gen concrete client
	g.Import(moduleImportPackage)
	unexportedClient := toLowerCamelCase(message.GoIdent) + "Client"
	g.P("type ", unexportedClient, " struct {")
	g.P("client ", moduleImportPackage.Ident("Client"))
	g.P("}")
	g.P()
	switch singleTon {
	case true:
		g.P("func (x *", unexportedClient, ") ", "Get() (*", message.GoIdent, ", error) {")
		g.P("_spfGenO := new(", message.GoIdent, ")")
		g.P("_spfGenErr := x.client.Get(", metaImportPackage.Ident("SingletonID"), ", _spfGenO)")
		g.P("if _spfGenErr != nil {")
		g.P("return nil, _spfGenErr")
		g.P("}")
		g.P("return _spfGenO, nil")
		g.P("}")
	case false:
		g.P("func (x *", unexportedClient, ") ", "Get(", primaryKey, " ", primaryKeyGoType, ") (*", message.GoIdent, ", error) {")
		g.P("_spfGenO := new(", message.GoIdent, ")")
		g.P("_spfGenID := ", metaImportPackage.Ident(metaIDConstructor[primaryKeyGoType]), "(", primaryKey, ")")
		g.P("_spfGenErr := x.client.Get(_spfGenID, _spfGenO)")
		g.P("if _spfGenErr != nil {")
		g.P("return nil, _spfGenErr")
		g.P("}")
		g.P("return _spfGenO, nil")
		g.P("}")
	}
	// gen create
	g.P()
	g.P("func (x *", unexportedClient, ") ", "Create(", toLowerCamelCase(message.GoIdent), " *", message.GoIdent, ") error {")
	g.P("return x.client.Create(", toLowerCamelCase(message.GoIdent), ")")
	g.P("}")
	g.P()
	// gen delete
	g.P("func (x *", unexportedClient, ") ", "Delete(", toLowerCamelCase(message.GoIdent), " *", message.GoIdent, ") error {")
	g.P("return x.client.Delete(", toLowerCamelCase(message.GoIdent), ")")
	g.P("}")
	g.P()
	// gen update
	g.P("func (x *", unexportedClient, ") ", "Update(", toLowerCamelCase(message.GoIdent), " *", message.GoIdent, ") error {")
	g.P("return x.client.Update(", toLowerCamelCase(message.GoIdent), ")")
	g.P("}")
	g.P()
}

func parseSaveInfo(m *protogen.Message) (bool, string, string, error) {
	opts := m.Desc.Options().(*descriptorpb.MessageOptions)
	// check if singleton
	isSingleton := proto.GetExtension(opts, modulegen.E_Singleton).(bool)
	if isSingleton {
		return true, "", "", nil
	}
	// if it's not singleton find the primary key
	primaryKey := proto.GetExtension(opts, modulegen.E_PrimaryKey).(string)
	if primaryKey == "" {
		return false, "", "", fmt.Errorf("%s has no primary key", m.Desc.Name())
	}
	fd := m.Desc.Fields().ByJSONName(primaryKey)
	if fd == nil {
		return false, "", "", fmt.Errorf("%s has no field named %s", m.Desc.Name(), primaryKey)
	}
	goType, exists := protoKindToGoType[fd.Kind()]
	if !exists {
		return false, "", "", fmt.Errorf("%s has unsupported primary key kind %s", m.Desc.Name(), fd.Kind())
	}
	return false, primaryKey, goType, nil
}

var protoKindToGoType = map[protoreflect.Kind]string{
	protoreflect.BoolKind:   "bool",
	protoreflect.StringKind: "string",
	protoreflect.BytesKind:  "[]byte",
	protoreflect.DoubleKind: "float64",
	protoreflect.Uint64Kind: "uint64",
	protoreflect.Uint32Kind: "uint32",
	protoreflect.Int32Kind:  "int32",
	protoreflect.Int64Kind:  "int64",
	// TODO support rest
}

var metaIDConstructor = map[string]string{
	"string": "NewStringID",
	"[]byte": "NewBytesID",
	// TODO add other types
}
