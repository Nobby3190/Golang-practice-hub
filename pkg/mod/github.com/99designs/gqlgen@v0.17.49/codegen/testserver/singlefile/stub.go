// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package singlefile

import (
	"context"

	introspection1 "github.com/99designs/gqlgen/codegen/testserver/singlefile/introspection"
	invalid_packagename "github.com/99designs/gqlgen/codegen/testserver/singlefile/invalid-packagename"
	"github.com/99designs/gqlgen/codegen/testserver/singlefile/otherpkg"
)

type Stub struct {
	BackedByInterfaceResolver struct {
		ID func(ctx context.Context, obj BackedByInterface) (string, error)
	}
	DeferModelResolver struct {
		Values func(ctx context.Context, obj *DeferModel) ([]string, error)
	}
	ErrorsResolver struct {
		A func(ctx context.Context, obj *Errors) (*Error, error)
		B func(ctx context.Context, obj *Errors) (*Error, error)
		C func(ctx context.Context, obj *Errors) (*Error, error)
		D func(ctx context.Context, obj *Errors) (*Error, error)
		E func(ctx context.Context, obj *Errors) (*Error, error)
	}
	ForcedResolverResolver struct {
		Field func(ctx context.Context, obj *ForcedResolver) (*Circle, error)
	}
	ModelMethodsResolver struct {
		ResolverField func(ctx context.Context, obj *ModelMethods) (bool, error)
	}
	MutationResolver struct {
		DefaultInput          func(ctx context.Context, input DefaultInput) (*DefaultParametersMirror, error)
		OverrideValueViaInput func(ctx context.Context, input FieldsOrderInput) (*FieldsOrderPayload, error)
		UpdateSomething       func(ctx context.Context, input SpecialInput) (string, error)
		UpdatePtrToPtr        func(ctx context.Context, input UpdatePtrToPtrOuter) (*PtrToPtrOuter, error)
	}
	OverlappingFieldsResolver struct {
		OldFoo func(ctx context.Context, obj *OverlappingFields) (int, error)
	}
	PanicsResolver struct {
		FieldScalarMarshal func(ctx context.Context, obj *Panics) ([]MarshalPanic, error)
		ArgUnmarshal       func(ctx context.Context, obj *Panics, u []MarshalPanic) (bool, error)
	}
	PetResolver struct {
		Friends func(ctx context.Context, obj *Pet, limit *int) ([]*Pet, error)
	}
	PrimitiveResolver struct {
		Value func(ctx context.Context, obj *Primitive) (int, error)
	}
	PrimitiveStringResolver struct {
		Value func(ctx context.Context, obj *PrimitiveString) (string, error)
		Len   func(ctx context.Context, obj *PrimitiveString) (int, error)
	}
	QueryResolver struct {
		InvalidIdentifier                func(ctx context.Context) (*invalid_packagename.InvalidIdentifier, error)
		Collision                        func(ctx context.Context) (*introspection1.It, error)
		MapInput                         func(ctx context.Context, input map[string]interface{}) (*bool, error)
		Recursive                        func(ctx context.Context, input *RecursiveInputSlice) (*bool, error)
		NestedInputs                     func(ctx context.Context, input [][]*OuterInput) (*bool, error)
		NestedOutputs                    func(ctx context.Context) ([][]*OuterObject, error)
		ModelMethods                     func(ctx context.Context) (*ModelMethods, error)
		User                             func(ctx context.Context, id int) (*User, error)
		NullableArg                      func(ctx context.Context, arg *int) (*string, error)
		InputSlice                       func(ctx context.Context, arg []string) (bool, error)
		InputNullableSlice               func(ctx context.Context, arg []string) (bool, error)
		InputOmittable                   func(ctx context.Context, arg OmittableInput) (string, error)
		ShapeUnion                       func(ctx context.Context) (ShapeUnion, error)
		Autobind                         func(ctx context.Context) (*Autobind, error)
		DeprecatedField                  func(ctx context.Context) (string, error)
		Overlapping                      func(ctx context.Context) (*OverlappingFields, error)
		DefaultParameters                func(ctx context.Context, falsyBoolean *bool, truthyBoolean *bool) (*DefaultParametersMirror, error)
		DeferCase1                       func(ctx context.Context) (*DeferModel, error)
		DeferCase2                       func(ctx context.Context) ([]*DeferModel, error)
		DirectiveArg                     func(ctx context.Context, arg string) (*string, error)
		DirectiveNullableArg             func(ctx context.Context, arg *int, arg2 *int, arg3 *string) (*string, error)
		DirectiveInputNullable           func(ctx context.Context, arg *InputDirectives) (*string, error)
		DirectiveInput                   func(ctx context.Context, arg InputDirectives) (*string, error)
		DirectiveInputType               func(ctx context.Context, arg InnerInput) (*string, error)
		DirectiveObject                  func(ctx context.Context) (*ObjectDirectives, error)
		DirectiveObjectWithCustomGoModel func(ctx context.Context) (*ObjectDirectivesWithCustomGoModel, error)
		DirectiveFieldDef                func(ctx context.Context, ret string) (string, error)
		DirectiveField                   func(ctx context.Context) (*string, error)
		DirectiveDouble                  func(ctx context.Context) (*string, error)
		DirectiveUnimplemented           func(ctx context.Context) (*string, error)
		EmbeddedCase1                    func(ctx context.Context) (*EmbeddedCase1, error)
		EmbeddedCase2                    func(ctx context.Context) (*EmbeddedCase2, error)
		EmbeddedCase3                    func(ctx context.Context) (*EmbeddedCase3, error)
		EnumInInput                      func(ctx context.Context, input *InputWithEnumValue) (EnumTest, error)
		Shapes                           func(ctx context.Context) ([]Shape, error)
		NoShape                          func(ctx context.Context) (Shape, error)
		Node                             func(ctx context.Context) (Node, error)
		NoShapeTypedNil                  func(ctx context.Context) (Shape, error)
		Animal                           func(ctx context.Context) (Animal, error)
		NotAnInterface                   func(ctx context.Context) (BackedByInterface, error)
		Dog                              func(ctx context.Context) (*Dog, error)
		Issue896a                        func(ctx context.Context) ([]*CheckIssue896, error)
		MapStringInterface               func(ctx context.Context, in map[string]interface{}) (map[string]interface{}, error)
		MapNestedStringInterface         func(ctx context.Context, in *NestedMapInput) (map[string]interface{}, error)
		ErrorBubble                      func(ctx context.Context) (*Error, error)
		ErrorBubbleList                  func(ctx context.Context) ([]*Error, error)
		ErrorList                        func(ctx context.Context) ([]*Error, error)
		Errors                           func(ctx context.Context) (*Errors, error)
		Valid                            func(ctx context.Context) (string, error)
		Invalid                          func(ctx context.Context) (string, error)
		Panics                           func(ctx context.Context) (*Panics, error)
		PrimitiveObject                  func(ctx context.Context) ([]Primitive, error)
		PrimitiveStringObject            func(ctx context.Context) ([]PrimitiveString, error)
		PtrToAnyContainer                func(ctx context.Context) (*PtrToAnyContainer, error)
		PtrToSliceContainer              func(ctx context.Context) (*PtrToSliceContainer, error)
		Infinity                         func(ctx context.Context) (float64, error)
		StringFromContextInterface       func(ctx context.Context) (*StringFromContextInterface, error)
		StringFromContextFunction        func(ctx context.Context) (string, error)
		DefaultScalar                    func(ctx context.Context, arg string) (string, error)
		Slices                           func(ctx context.Context) (*Slices, error)
		ScalarSlice                      func(ctx context.Context) ([]byte, error)
		Fallback                         func(ctx context.Context, arg FallbackToStringEncoding) (FallbackToStringEncoding, error)
		OptionalUnion                    func(ctx context.Context) (TestUnion, error)
		VOkCaseValue                     func(ctx context.Context) (*VOkCaseValue, error)
		VOkCaseNil                       func(ctx context.Context) (*VOkCaseNil, error)
		ValidType                        func(ctx context.Context) (*ValidType, error)
		VariadicModel                    func(ctx context.Context) (*VariadicModel, error)
		WrappedStruct                    func(ctx context.Context) (*WrappedStruct, error)
		WrappedScalar                    func(ctx context.Context) (otherpkg.Scalar, error)
		WrappedMap                       func(ctx context.Context) (WrappedMap, error)
		WrappedSlice                     func(ctx context.Context) (WrappedSlice, error)
	}
	SubscriptionResolver struct {
		Updated                func(ctx context.Context) (<-chan string, error)
		InitPayload            func(ctx context.Context) (<-chan string, error)
		DirectiveArg           func(ctx context.Context, arg string) (<-chan *string, error)
		DirectiveNullableArg   func(ctx context.Context, arg *int, arg2 *int, arg3 *string) (<-chan *string, error)
		DirectiveDouble        func(ctx context.Context) (<-chan *string, error)
		DirectiveUnimplemented func(ctx context.Context) (<-chan *string, error)
		Issue896b              func(ctx context.Context) (<-chan []*CheckIssue896, error)
		ErrorRequired          func(ctx context.Context) (<-chan *Error, error)
	}
	UserResolver struct {
		Friends func(ctx context.Context, obj *User) ([]*User, error)
		Pets    func(ctx context.Context, obj *User, limit *int) ([]*Pet, error)
	}
	WrappedMapResolver struct {
		Get func(ctx context.Context, obj WrappedMap, key string) (string, error)
	}
	WrappedSliceResolver struct {
		Get func(ctx context.Context, obj WrappedSlice, idx int) (string, error)
	}

	FieldsOrderInputResolver struct {
		OverrideFirstField func(ctx context.Context, obj *FieldsOrderInput, data *string) error
	}
}

func (r *Stub) BackedByInterface() BackedByInterfaceResolver {
	return &stubBackedByInterface{r}
}
func (r *Stub) DeferModel() DeferModelResolver {
	return &stubDeferModel{r}
}
func (r *Stub) Errors() ErrorsResolver {
	return &stubErrors{r}
}
func (r *Stub) ForcedResolver() ForcedResolverResolver {
	return &stubForcedResolver{r}
}
func (r *Stub) ModelMethods() ModelMethodsResolver {
	return &stubModelMethods{r}
}
func (r *Stub) Mutation() MutationResolver {
	return &stubMutation{r}
}
func (r *Stub) OverlappingFields() OverlappingFieldsResolver {
	return &stubOverlappingFields{r}
}
func (r *Stub) Panics() PanicsResolver {
	return &stubPanics{r}
}
func (r *Stub) Pet() PetResolver {
	return &stubPet{r}
}
func (r *Stub) Primitive() PrimitiveResolver {
	return &stubPrimitive{r}
}
func (r *Stub) PrimitiveString() PrimitiveStringResolver {
	return &stubPrimitiveString{r}
}
func (r *Stub) Query() QueryResolver {
	return &stubQuery{r}
}
func (r *Stub) Subscription() SubscriptionResolver {
	return &stubSubscription{r}
}
func (r *Stub) User() UserResolver {
	return &stubUser{r}
}
func (r *Stub) WrappedMap() WrappedMapResolver {
	return &stubWrappedMap{r}
}
func (r *Stub) WrappedSlice() WrappedSliceResolver {
	return &stubWrappedSlice{r}
}

func (r *Stub) FieldsOrderInput() FieldsOrderInputResolver {
	return &stubFieldsOrderInput{r}
}

type stubBackedByInterface struct{ *Stub }

func (r *stubBackedByInterface) ID(ctx context.Context, obj BackedByInterface) (string, error) {
	return r.BackedByInterfaceResolver.ID(ctx, obj)
}

type stubDeferModel struct{ *Stub }

func (r *stubDeferModel) Values(ctx context.Context, obj *DeferModel) ([]string, error) {
	return r.DeferModelResolver.Values(ctx, obj)
}

type stubErrors struct{ *Stub }

func (r *stubErrors) A(ctx context.Context, obj *Errors) (*Error, error) {
	return r.ErrorsResolver.A(ctx, obj)
}
func (r *stubErrors) B(ctx context.Context, obj *Errors) (*Error, error) {
	return r.ErrorsResolver.B(ctx, obj)
}
func (r *stubErrors) C(ctx context.Context, obj *Errors) (*Error, error) {
	return r.ErrorsResolver.C(ctx, obj)
}
func (r *stubErrors) D(ctx context.Context, obj *Errors) (*Error, error) {
	return r.ErrorsResolver.D(ctx, obj)
}
func (r *stubErrors) E(ctx context.Context, obj *Errors) (*Error, error) {
	return r.ErrorsResolver.E(ctx, obj)
}

type stubForcedResolver struct{ *Stub }

func (r *stubForcedResolver) Field(ctx context.Context, obj *ForcedResolver) (*Circle, error) {
	return r.ForcedResolverResolver.Field(ctx, obj)
}

type stubModelMethods struct{ *Stub }

func (r *stubModelMethods) ResolverField(ctx context.Context, obj *ModelMethods) (bool, error) {
	return r.ModelMethodsResolver.ResolverField(ctx, obj)
}

type stubMutation struct{ *Stub }

func (r *stubMutation) DefaultInput(ctx context.Context, input DefaultInput) (*DefaultParametersMirror, error) {
	return r.MutationResolver.DefaultInput(ctx, input)
}
func (r *stubMutation) OverrideValueViaInput(ctx context.Context, input FieldsOrderInput) (*FieldsOrderPayload, error) {
	return r.MutationResolver.OverrideValueViaInput(ctx, input)
}
func (r *stubMutation) UpdateSomething(ctx context.Context, input SpecialInput) (string, error) {
	return r.MutationResolver.UpdateSomething(ctx, input)
}
func (r *stubMutation) UpdatePtrToPtr(ctx context.Context, input UpdatePtrToPtrOuter) (*PtrToPtrOuter, error) {
	return r.MutationResolver.UpdatePtrToPtr(ctx, input)
}

type stubOverlappingFields struct{ *Stub }

func (r *stubOverlappingFields) OldFoo(ctx context.Context, obj *OverlappingFields) (int, error) {
	return r.OverlappingFieldsResolver.OldFoo(ctx, obj)
}

type stubPanics struct{ *Stub }

func (r *stubPanics) FieldScalarMarshal(ctx context.Context, obj *Panics) ([]MarshalPanic, error) {
	return r.PanicsResolver.FieldScalarMarshal(ctx, obj)
}
func (r *stubPanics) ArgUnmarshal(ctx context.Context, obj *Panics, u []MarshalPanic) (bool, error) {
	return r.PanicsResolver.ArgUnmarshal(ctx, obj, u)
}

type stubPet struct{ *Stub }

func (r *stubPet) Friends(ctx context.Context, obj *Pet, limit *int) ([]*Pet, error) {
	return r.PetResolver.Friends(ctx, obj, limit)
}

type stubPrimitive struct{ *Stub }

func (r *stubPrimitive) Value(ctx context.Context, obj *Primitive) (int, error) {
	return r.PrimitiveResolver.Value(ctx, obj)
}

type stubPrimitiveString struct{ *Stub }

func (r *stubPrimitiveString) Value(ctx context.Context, obj *PrimitiveString) (string, error) {
	return r.PrimitiveStringResolver.Value(ctx, obj)
}
func (r *stubPrimitiveString) Len(ctx context.Context, obj *PrimitiveString) (int, error) {
	return r.PrimitiveStringResolver.Len(ctx, obj)
}

type stubQuery struct{ *Stub }

func (r *stubQuery) InvalidIdentifier(ctx context.Context) (*invalid_packagename.InvalidIdentifier, error) {
	return r.QueryResolver.InvalidIdentifier(ctx)
}
func (r *stubQuery) Collision(ctx context.Context) (*introspection1.It, error) {
	return r.QueryResolver.Collision(ctx)
}
func (r *stubQuery) MapInput(ctx context.Context, input map[string]interface{}) (*bool, error) {
	return r.QueryResolver.MapInput(ctx, input)
}
func (r *stubQuery) Recursive(ctx context.Context, input *RecursiveInputSlice) (*bool, error) {
	return r.QueryResolver.Recursive(ctx, input)
}
func (r *stubQuery) NestedInputs(ctx context.Context, input [][]*OuterInput) (*bool, error) {
	return r.QueryResolver.NestedInputs(ctx, input)
}
func (r *stubQuery) NestedOutputs(ctx context.Context) ([][]*OuterObject, error) {
	return r.QueryResolver.NestedOutputs(ctx)
}
func (r *stubQuery) ModelMethods(ctx context.Context) (*ModelMethods, error) {
	return r.QueryResolver.ModelMethods(ctx)
}
func (r *stubQuery) User(ctx context.Context, id int) (*User, error) {
	return r.QueryResolver.User(ctx, id)
}
func (r *stubQuery) NullableArg(ctx context.Context, arg *int) (*string, error) {
	return r.QueryResolver.NullableArg(ctx, arg)
}
func (r *stubQuery) InputSlice(ctx context.Context, arg []string) (bool, error) {
	return r.QueryResolver.InputSlice(ctx, arg)
}
func (r *stubQuery) InputNullableSlice(ctx context.Context, arg []string) (bool, error) {
	return r.QueryResolver.InputNullableSlice(ctx, arg)
}
func (r *stubQuery) InputOmittable(ctx context.Context, arg OmittableInput) (string, error) {
	return r.QueryResolver.InputOmittable(ctx, arg)
}
func (r *stubQuery) ShapeUnion(ctx context.Context) (ShapeUnion, error) {
	return r.QueryResolver.ShapeUnion(ctx)
}
func (r *stubQuery) Autobind(ctx context.Context) (*Autobind, error) {
	return r.QueryResolver.Autobind(ctx)
}
func (r *stubQuery) DeprecatedField(ctx context.Context) (string, error) {
	return r.QueryResolver.DeprecatedField(ctx)
}
func (r *stubQuery) Overlapping(ctx context.Context) (*OverlappingFields, error) {
	return r.QueryResolver.Overlapping(ctx)
}
func (r *stubQuery) DefaultParameters(ctx context.Context, falsyBoolean *bool, truthyBoolean *bool) (*DefaultParametersMirror, error) {
	return r.QueryResolver.DefaultParameters(ctx, falsyBoolean, truthyBoolean)
}
func (r *stubQuery) DeferCase1(ctx context.Context) (*DeferModel, error) {
	return r.QueryResolver.DeferCase1(ctx)
}
func (r *stubQuery) DeferCase2(ctx context.Context) ([]*DeferModel, error) {
	return r.QueryResolver.DeferCase2(ctx)
}
func (r *stubQuery) DirectiveArg(ctx context.Context, arg string) (*string, error) {
	return r.QueryResolver.DirectiveArg(ctx, arg)
}
func (r *stubQuery) DirectiveNullableArg(ctx context.Context, arg *int, arg2 *int, arg3 *string) (*string, error) {
	return r.QueryResolver.DirectiveNullableArg(ctx, arg, arg2, arg3)
}
func (r *stubQuery) DirectiveInputNullable(ctx context.Context, arg *InputDirectives) (*string, error) {
	return r.QueryResolver.DirectiveInputNullable(ctx, arg)
}
func (r *stubQuery) DirectiveInput(ctx context.Context, arg InputDirectives) (*string, error) {
	return r.QueryResolver.DirectiveInput(ctx, arg)
}
func (r *stubQuery) DirectiveInputType(ctx context.Context, arg InnerInput) (*string, error) {
	return r.QueryResolver.DirectiveInputType(ctx, arg)
}
func (r *stubQuery) DirectiveObject(ctx context.Context) (*ObjectDirectives, error) {
	return r.QueryResolver.DirectiveObject(ctx)
}
func (r *stubQuery) DirectiveObjectWithCustomGoModel(ctx context.Context) (*ObjectDirectivesWithCustomGoModel, error) {
	return r.QueryResolver.DirectiveObjectWithCustomGoModel(ctx)
}
func (r *stubQuery) DirectiveFieldDef(ctx context.Context, ret string) (string, error) {
	return r.QueryResolver.DirectiveFieldDef(ctx, ret)
}
func (r *stubQuery) DirectiveField(ctx context.Context) (*string, error) {
	return r.QueryResolver.DirectiveField(ctx)
}
func (r *stubQuery) DirectiveDouble(ctx context.Context) (*string, error) {
	return r.QueryResolver.DirectiveDouble(ctx)
}
func (r *stubQuery) DirectiveUnimplemented(ctx context.Context) (*string, error) {
	return r.QueryResolver.DirectiveUnimplemented(ctx)
}
func (r *stubQuery) EmbeddedCase1(ctx context.Context) (*EmbeddedCase1, error) {
	return r.QueryResolver.EmbeddedCase1(ctx)
}
func (r *stubQuery) EmbeddedCase2(ctx context.Context) (*EmbeddedCase2, error) {
	return r.QueryResolver.EmbeddedCase2(ctx)
}
func (r *stubQuery) EmbeddedCase3(ctx context.Context) (*EmbeddedCase3, error) {
	return r.QueryResolver.EmbeddedCase3(ctx)
}
func (r *stubQuery) EnumInInput(ctx context.Context, input *InputWithEnumValue) (EnumTest, error) {
	return r.QueryResolver.EnumInInput(ctx, input)
}
func (r *stubQuery) Shapes(ctx context.Context) ([]Shape, error) {
	return r.QueryResolver.Shapes(ctx)
}
func (r *stubQuery) NoShape(ctx context.Context) (Shape, error) {
	return r.QueryResolver.NoShape(ctx)
}
func (r *stubQuery) Node(ctx context.Context) (Node, error) {
	return r.QueryResolver.Node(ctx)
}
func (r *stubQuery) NoShapeTypedNil(ctx context.Context) (Shape, error) {
	return r.QueryResolver.NoShapeTypedNil(ctx)
}
func (r *stubQuery) Animal(ctx context.Context) (Animal, error) {
	return r.QueryResolver.Animal(ctx)
}
func (r *stubQuery) NotAnInterface(ctx context.Context) (BackedByInterface, error) {
	return r.QueryResolver.NotAnInterface(ctx)
}
func (r *stubQuery) Dog(ctx context.Context) (*Dog, error) {
	return r.QueryResolver.Dog(ctx)
}
func (r *stubQuery) Issue896a(ctx context.Context) ([]*CheckIssue896, error) {
	return r.QueryResolver.Issue896a(ctx)
}
func (r *stubQuery) MapStringInterface(ctx context.Context, in map[string]interface{}) (map[string]interface{}, error) {
	return r.QueryResolver.MapStringInterface(ctx, in)
}
func (r *stubQuery) MapNestedStringInterface(ctx context.Context, in *NestedMapInput) (map[string]interface{}, error) {
	return r.QueryResolver.MapNestedStringInterface(ctx, in)
}
func (r *stubQuery) ErrorBubble(ctx context.Context) (*Error, error) {
	return r.QueryResolver.ErrorBubble(ctx)
}
func (r *stubQuery) ErrorBubbleList(ctx context.Context) ([]*Error, error) {
	return r.QueryResolver.ErrorBubbleList(ctx)
}
func (r *stubQuery) ErrorList(ctx context.Context) ([]*Error, error) {
	return r.QueryResolver.ErrorList(ctx)
}
func (r *stubQuery) Errors(ctx context.Context) (*Errors, error) {
	return r.QueryResolver.Errors(ctx)
}
func (r *stubQuery) Valid(ctx context.Context) (string, error) {
	return r.QueryResolver.Valid(ctx)
}
func (r *stubQuery) Invalid(ctx context.Context) (string, error) {
	return r.QueryResolver.Invalid(ctx)
}
func (r *stubQuery) Panics(ctx context.Context) (*Panics, error) {
	return r.QueryResolver.Panics(ctx)
}
func (r *stubQuery) PrimitiveObject(ctx context.Context) ([]Primitive, error) {
	return r.QueryResolver.PrimitiveObject(ctx)
}
func (r *stubQuery) PrimitiveStringObject(ctx context.Context) ([]PrimitiveString, error) {
	return r.QueryResolver.PrimitiveStringObject(ctx)
}
func (r *stubQuery) PtrToAnyContainer(ctx context.Context) (*PtrToAnyContainer, error) {
	return r.QueryResolver.PtrToAnyContainer(ctx)
}
func (r *stubQuery) PtrToSliceContainer(ctx context.Context) (*PtrToSliceContainer, error) {
	return r.QueryResolver.PtrToSliceContainer(ctx)
}
func (r *stubQuery) Infinity(ctx context.Context) (float64, error) {
	return r.QueryResolver.Infinity(ctx)
}
func (r *stubQuery) StringFromContextInterface(ctx context.Context) (*StringFromContextInterface, error) {
	return r.QueryResolver.StringFromContextInterface(ctx)
}
func (r *stubQuery) StringFromContextFunction(ctx context.Context) (string, error) {
	return r.QueryResolver.StringFromContextFunction(ctx)
}
func (r *stubQuery) DefaultScalar(ctx context.Context, arg string) (string, error) {
	return r.QueryResolver.DefaultScalar(ctx, arg)
}
func (r *stubQuery) Slices(ctx context.Context) (*Slices, error) {
	return r.QueryResolver.Slices(ctx)
}
func (r *stubQuery) ScalarSlice(ctx context.Context) ([]byte, error) {
	return r.QueryResolver.ScalarSlice(ctx)
}
func (r *stubQuery) Fallback(ctx context.Context, arg FallbackToStringEncoding) (FallbackToStringEncoding, error) {
	return r.QueryResolver.Fallback(ctx, arg)
}
func (r *stubQuery) OptionalUnion(ctx context.Context) (TestUnion, error) {
	return r.QueryResolver.OptionalUnion(ctx)
}
func (r *stubQuery) VOkCaseValue(ctx context.Context) (*VOkCaseValue, error) {
	return r.QueryResolver.VOkCaseValue(ctx)
}
func (r *stubQuery) VOkCaseNil(ctx context.Context) (*VOkCaseNil, error) {
	return r.QueryResolver.VOkCaseNil(ctx)
}
func (r *stubQuery) ValidType(ctx context.Context) (*ValidType, error) {
	return r.QueryResolver.ValidType(ctx)
}
func (r *stubQuery) VariadicModel(ctx context.Context) (*VariadicModel, error) {
	return r.QueryResolver.VariadicModel(ctx)
}
func (r *stubQuery) WrappedStruct(ctx context.Context) (*WrappedStruct, error) {
	return r.QueryResolver.WrappedStruct(ctx)
}
func (r *stubQuery) WrappedScalar(ctx context.Context) (otherpkg.Scalar, error) {
	return r.QueryResolver.WrappedScalar(ctx)
}
func (r *stubQuery) WrappedMap(ctx context.Context) (WrappedMap, error) {
	return r.QueryResolver.WrappedMap(ctx)
}
func (r *stubQuery) WrappedSlice(ctx context.Context) (WrappedSlice, error) {
	return r.QueryResolver.WrappedSlice(ctx)
}

type stubSubscription struct{ *Stub }

func (r *stubSubscription) Updated(ctx context.Context) (<-chan string, error) {
	return r.SubscriptionResolver.Updated(ctx)
}
func (r *stubSubscription) InitPayload(ctx context.Context) (<-chan string, error) {
	return r.SubscriptionResolver.InitPayload(ctx)
}
func (r *stubSubscription) DirectiveArg(ctx context.Context, arg string) (<-chan *string, error) {
	return r.SubscriptionResolver.DirectiveArg(ctx, arg)
}
func (r *stubSubscription) DirectiveNullableArg(ctx context.Context, arg *int, arg2 *int, arg3 *string) (<-chan *string, error) {
	return r.SubscriptionResolver.DirectiveNullableArg(ctx, arg, arg2, arg3)
}
func (r *stubSubscription) DirectiveDouble(ctx context.Context) (<-chan *string, error) {
	return r.SubscriptionResolver.DirectiveDouble(ctx)
}
func (r *stubSubscription) DirectiveUnimplemented(ctx context.Context) (<-chan *string, error) {
	return r.SubscriptionResolver.DirectiveUnimplemented(ctx)
}
func (r *stubSubscription) Issue896b(ctx context.Context) (<-chan []*CheckIssue896, error) {
	return r.SubscriptionResolver.Issue896b(ctx)
}
func (r *stubSubscription) ErrorRequired(ctx context.Context) (<-chan *Error, error) {
	return r.SubscriptionResolver.ErrorRequired(ctx)
}

type stubUser struct{ *Stub }

func (r *stubUser) Friends(ctx context.Context, obj *User) ([]*User, error) {
	return r.UserResolver.Friends(ctx, obj)
}
func (r *stubUser) Pets(ctx context.Context, obj *User, limit *int) ([]*Pet, error) {
	return r.UserResolver.Pets(ctx, obj, limit)
}

type stubWrappedMap struct{ *Stub }

func (r *stubWrappedMap) Get(ctx context.Context, obj WrappedMap, key string) (string, error) {
	return r.WrappedMapResolver.Get(ctx, obj, key)
}

type stubWrappedSlice struct{ *Stub }

func (r *stubWrappedSlice) Get(ctx context.Context, obj WrappedSlice, idx int) (string, error) {
	return r.WrappedSliceResolver.Get(ctx, obj, idx)
}

type stubFieldsOrderInput struct{ *Stub }

func (r *stubFieldsOrderInput) OverrideFirstField(ctx context.Context, obj *FieldsOrderInput, data *string) error {
	return r.FieldsOrderInputResolver.OverrideFirstField(ctx, obj, data)
}
