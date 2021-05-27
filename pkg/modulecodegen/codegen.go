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
const clientImportPackage = protogen.GoImportPath("github.com/fdymylja/tmos/runtime/client")

const GenCodeFileSuffix = ".starport.go"

func PluginRunner(plugin *protogen.Plugin) error {
	groups := make(map[protogen.GoImportPath][]*protogen.File)
	for _, f := range plugin.Files {
		groups[f.GoImportPath] = append(groups[f.GoImportPath], f)
	}
	// then we parse single groups
	for _, group := range groups {
		err := genFiles(plugin, group)
		if err != nil {
			return err
		}
	}
	return nil
}

func genFiles(gen *protogen.Plugin, group []*protogen.File) error {
	for _, file := range group {
		genFile(file, gen)
	}
	return nil
}

func genFile(file *protogen.File, gen *protogen.Plugin) {

	if !meetsRequirements(file) {
		return
	}

	filename := fmt.Sprintf("%s%s", file.GeneratedFilenamePrefix, GenCodeFileSuffix)
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
	// gen schema
	for _, obj := range stateObjects {
		err := genSchema(objectsFile, obj)
		if err != nil {
			gen.Error(err)
		}
	}
	// gen clientset
	genClientSet(objectsFile, stateObjects, stateTransitions)
}

func genClientSet(g *protogen.GeneratedFile, objects []*protogen.Message, transitions []*protogen.Message) {
	// gen clientset interface
	g.P("type ClientSet interface {")
	// add state objects client interface
	for _, obj := range objects {
		// if it ends with s we don't add the 's' to indicate the plural of types
		switch strings.HasSuffix(obj.GoIdent.GoName, "s") || isSingletonObject(obj) {
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

	// gen constructor
	g.P("func NewClientSet(client ", clientImportPackage.Ident("RuntimeClient"), ") ClientSet {")
	g.P("return &clientSet{")
	g.P("client: client", ",") // the normal module client
	// add other clients
	for _, obj := range objects {
		unexportedClient := toLowerCamelCase(obj.GoIdent) + "Client"
		g.P(unexportedClient, ": ", "&", unexportedClient, "{client: client}", ",")
	}
	g.P("}")

	g.P("}")
	g.P()
	// gen client type
	g.P("type clientSet struct {")
	g.P("client ", clientImportPackage.Ident("RuntimeClient"))
	// include state objects clients
	for _, obj := range objects {
		unexportedClient := toLowerCamelCase(obj.GoIdent) + "Client"
		exportedClient := obj.GoIdent.GoName + "Client"
		g.P("// ", unexportedClient, " is the client used to interact with ", obj.GoIdent)
		g.P(unexportedClient, " ", exportedClient)
	}
	g.P("}")
	g.P()
	// gen client concrete methods
	for _, obj := range objects {
		unexportedClient := toLowerCamelCase(obj.GoIdent) + "Client"
		// if it ends with s we don't add the 's' to indicate the plural of types
		switch strings.HasSuffix(obj.GoIdent.GoName, "s") || isSingletonObject(obj) {
		case false:
			g.P("func (x *clientSet) ", obj.GoIdent, "s()", " ", obj.GoIdent, "Client", " {")
			g.P("return x.", unexportedClient)
			g.P("}")
			g.P()
		case true:
			g.P("func (x *clientSet) ", obj.GoIdent, "()", " ", obj.GoIdent, "Client", " {")
			g.P("return x.", unexportedClient)
			g.P("}")
			g.P()
		}
	}

	// gen state transitions interface
	for _, t := range transitions {
		g.P("func (x *clientSet) Exec", t.GoIdent.GoName, "(msg *", t.GoIdent.GoName, ") error {")
		g.P("return x.client.Deliver(msg)")
		g.P("}")
		g.P()
	}
	g.P()
}

func isSingletonObject(obj *protogen.Message) bool {
	opts := obj.Desc.Options().(*descriptorpb.MessageOptions)
	return proto.GetExtension(opts, modulegen.E_Singleton).(bool)
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

	genStateObjectClient(g, message)
}

func genStateObjectClient(g *protogen.GeneratedFile, message *protogen.Message) {
	singleTon, primaryKey, primaryKeyGoType, err := parseSaveInfo(message)
	if err != nil {
		panic(err) // TODO not with panic
	}
	exportedClient := message.GoIdent.GoName + "Client"
	g.P("type ", exportedClient, " interface {")
	switch singleTon {
	case true:
		g.P("Get(opts ...", clientImportPackage.Ident("GetOption"), ") (*", message.GoIdent, ", error)")
	case false:
		g.P("Get(", primaryKey, " ", primaryKeyGoType, ", opts ...", clientImportPackage.Ident("GetOption"), ") (*", message.GoIdent, ", error)")
	}
	g.P("Create(", toLowerCamelCase(message.GoIdent), " *", message.GoIdent, ", opts ...", clientImportPackage.Ident("CreateOption"), ") error")
	g.P("Delete(", toLowerCamelCase(message.GoIdent), " *", message.GoIdent, ", opts ...", clientImportPackage.Ident("DeleteOption"), ") error")
	g.P("Update(", toLowerCamelCase(message.GoIdent), " *", message.GoIdent, ", opts ...", clientImportPackage.Ident("UpdateOption"), ") error")
	g.P("}")
	g.P()

	unexportedClient := toLowerCamelCase(message.GoIdent) + "Client" // is the concrete client name
	// gen concrete client
	g.Import(clientImportPackage)
	g.P("type ", unexportedClient, " struct {")
	g.P("client ", clientImportPackage.Ident("RuntimeClient"))
	g.P("}")
	g.P()
	switch singleTon {
	case true:
		g.P("func (x *", unexportedClient, ") ", "Get(opts ...", clientImportPackage.Ident("GetOption"), ") (*", message.GoIdent, ", error) {")
		g.P("_spfGenO := new(", message.GoIdent, ")")
		g.P("_spfGenErr := x.client.Get(", metaImportPackage.Ident("SingletonID"), ", _spfGenO, opts...)")
		g.P("if _spfGenErr != nil {")
		g.P("return nil, _spfGenErr")
		g.P("}")
		g.P("return _spfGenO, nil")
		g.P("}")
	case false:
		g.P("func (x *", unexportedClient, ") ", "Get(", primaryKey, " ", primaryKeyGoType, ", opts... ", clientImportPackage.Ident("GetOption"), ") (*", message.GoIdent, ", error) {")
		g.P("_spfGenO := new(", message.GoIdent, ")")
		g.P("_spfGenID := ", metaImportPackage.Ident(metaIDConstructor[primaryKeyGoType]), "(", primaryKey, ")")
		g.P("_spfGenErr := x.client.Get(_spfGenID, _spfGenO, opts...)")
		g.P("if _spfGenErr != nil {")
		g.P("return nil, _spfGenErr")
		g.P("}")
		g.P("return _spfGenO, nil")
		g.P("}")
	}
	// gen create
	g.P()
	g.P("func (x *", unexportedClient, ") ", "Create(", toLowerCamelCase(message.GoIdent), " *", message.GoIdent, ", opts ...", clientImportPackage.Ident("CreateOption"), ") error {")
	g.P("return x.client.Create(", toLowerCamelCase(message.GoIdent), ", opts...)")
	g.P("}")
	g.P()
	// gen delete
	g.P("func (x *", unexportedClient, ") ", "Delete(", toLowerCamelCase(message.GoIdent), " *", message.GoIdent, ", opts ...", clientImportPackage.Ident("DeleteOption"), ") error {")
	g.P("return x.client.Delete(", toLowerCamelCase(message.GoIdent), ", opts...)")
	g.P("}")
	g.P()
	// gen update
	g.P("func (x *", unexportedClient, ") ", "Update(", toLowerCamelCase(message.GoIdent), " *", message.GoIdent, ", opts ... ", clientImportPackage.Ident("UpdateOption"), ") error {")
	g.P("return x.client.Update(", toLowerCamelCase(message.GoIdent), ", opts...)")
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
