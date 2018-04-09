package handlers

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/gobuffalo/uuid"

	"github.com/transcom/mymove/pkg/auth/context"
	servicememberop "github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/service_members"
	"github.com/transcom/mymove/pkg/gen/internalmessages"
	"github.com/transcom/mymove/pkg/models"
)

func Test_ServiceMember(t *testing.T) {
	fmt.Println("testing233")
}

func (suite *HandlerSuite) TestShowServiceMemberHandler() {
	t := suite.T()

	// Given: A servicemember and a user
	user := models.User{
		LoginGovUUID:  uuid.Must(uuid.NewV4()),
		LoginGovEmail: "email@example.com",
		Type:          internalmessages.UserTypeUNKNOWN,
	}
	suite.mustSave(&user)

	newServiceMember := models.ServiceMember{
		UserID: user.ID,
	}
	suite.mustSave(&newServiceMember)

	// And: the context contains the auth values
	req := httptest.NewRequest("GET", "/service_members/some_id", nil)
	ctx := req.Context()
	ctx = context.PopulateAuthContext(ctx, user.ID, "fake token")
	req = req.WithContext(ctx)

	params := servicememberop.ShowServiceMemberParams{
		HTTPRequest:     req,
		ServiceMemberID: strfmt.UUID(newServiceMember.ID.String()),
	}
	// And: show ServiceMember is queried
	showHandler := ShowServiceMemberHandler(NewHandlerContext(suite.db, suite.logger))
	showResponse := showHandler.Handle(params)

	// Then: Expect a 200 status code
	okResponse := showResponse.(*servicememberop.ShowServiceMemberOK)
	servicemember := okResponse.Payload

	// And: Returned query to include our added servicemember
	if servicemember.UserID.String() != user.ID.String() {
		t.Errorf("Expected an servicemember to have user ID '%v'. None do.", user.ID)
	}

}

func (suite *HandlerSuite) TestShowServiceMemberHandlerNoUser() {
	t := suite.T()

	// Given: A servicemember with a user that isn't logged in
	user := models.User{
		LoginGovUUID:  uuid.Must(uuid.NewV4()),
		LoginGovEmail: "email@example.com",
		Type:          internalmessages.UserTypeUNKNOWN,
	}
	suite.mustSave(&user)

	servicemember := models.ServiceMember{
		UserID: user.ID,
	}
	suite.mustSave(&servicemember)

	req := httptest.NewRequest("GET", "/service_members/some_id", nil)
	showServiceMemberParams := servicememberop.NewShowServiceMemberParams()
	showServiceMemberParams.HTTPRequest = req

	// And: Show servicemember is queried
	showHandler := ShowServiceMemberHandler(NewHandlerContext(suite.db, suite.logger))
	showResponse := showHandler.Handle(showServiceMemberParams)

	// Then: Expect a 401 unauthorized
	_, ok := showResponse.(*servicememberop.ShowServiceMemberUnauthorized)
	if !ok {
		t.Errorf("Expected to get an unauthorized response, but got something else.")
	}
}

func (suite *HandlerSuite) TestShowServiceMemberWrongUser() {
	t := suite.T()

	// Given: A servicemember with a not-logged-in user and a separate logged-in user
	notLoggedInUser := models.User{
		LoginGovUUID:  uuid.Must(uuid.NewV4()),
		LoginGovEmail: "email@example.com",
		Type:          internalmessages.UserTypeUNKNOWN,
	}
	suite.mustSave(&notLoggedInUser)

	loggedInUser := models.User{
		LoginGovUUID:  uuid.Must(uuid.NewV4()),
		LoginGovEmail: "email2@example.com",
		Type:          internalmessages.UserTypeUNKNOWN,
	}
	suite.mustSave(&loggedInUser)

	// When: A servicemember is created for not-logged-in-user
	newServiceMember := models.ServiceMember{
		UserID: notLoggedInUser.ID,
	}
	suite.mustSave(&newServiceMember)

	// And: the context contains the auth values for logged-in user
	req := httptest.NewRequest("GET", "/service_members/some_id", nil)
	ctx := req.Context()
	ctx = context.PopulateAuthContext(ctx, loggedInUser.ID, "fake token")
	req = req.WithContext(ctx)
	showServiceMemberParams := servicememberop.ShowServiceMemberParams{
		HTTPRequest:     req,
		ServiceMemberID: strfmt.UUID(newServiceMember.ID.String()),
	}
	// And: Show servicemember is queried
	showHandler := ShowServiceMemberHandler(NewHandlerContext(suite.db, suite.logger))
	showResponse := showHandler.Handle(showServiceMemberParams)

	_, ok := showResponse.(*servicememberop.ShowServiceMemberForbidden)
	if !ok {
		t.Fatalf("Request failed: %#v", showResponse)
	}
}

func (suite *HandlerSuite) TestSubmitServiceMemberHandlerAllValues() {
	t := suite.T()

	// Given: A logged-in user
	user := models.User{
		LoginGovUUID:  uuid.Must(uuid.NewV4()),
		LoginGovEmail: "email@example.com",
		Type:          internalmessages.UserTypeUNKNOWN,
	}
	suite.mustSave(&user)

	// When: a new ServiceMember is posted
	newServiceMemberPayload := internalmessages.CreateServiceMemberPayload{
		UserID:                    strfmt.UUID(user.ID.String()),
		Edipi:                     swag.String("random string bla"),
		FirstName:                 swag.String("random string bla"),
		MiddleInitial:             swag.String("random string bla"),
		LastName:                  swag.String("random string bla"),
		Suffix:                    swag.String("random string bla"),
		Telephone:                 swag.String("random string bla"),
		SecondaryTelephone:        swag.String("random string bla"),
		PersonalEmail:             swag.String("random string bla"),
		PhoneIsPreferred:          swag.Bool(false),
		SecondaryPhoneIsPreferred: swag.Bool(false),
		EmailIsPreferred:          swag.Bool(true),
		ResidentialAddress:        fakeAddress(),
		BackupMailingAddress:      fakeAddress(),
	}

	req := httptest.NewRequest("GET", "/service_members/some_id", nil)
	params := servicememberop.CreateServiceMemberParams{
		CreateServiceMemberPayload: &newServiceMemberPayload,
		HTTPRequest:                req,
	}

	// And: the context contains the auth values for logged-in user
	ctx := params.HTTPRequest.Context()
	ctx = context.PopulateAuthContext(ctx, user.ID, "fake token")
	params.HTTPRequest = params.HTTPRequest.WithContext(ctx)

	handler := CreateServiceMemberHandler(NewHandlerContext(suite.db, suite.logger))
	response := handler.Handle(params)

	_, ok := response.(*servicememberop.CreateServiceMemberCreated)
	if !ok {
		t.Fatalf("Request failed: %#v", response)
	}

	// Then: we expect a servicemember to have been created for the user
	query := suite.db.Where(fmt.Sprintf("user_id='%v'", user.ID))
	servicemembers := []models.ServiceMember{}
	query.All(&servicemembers)

	if len(servicemembers) != 1 {
		t.Errorf("Expected to find 1 servicemember but found %v", len(servicemembers))
	}
}

func (suite *HandlerSuite) TestCreateServiceMemberHandlerNoUserID() {
	t := suite.T()
	// Given: no authentication values in context
	// When: a new ServiceMember is posted
	servicememberPayload := internalmessages.CreateServiceMemberPayload{}
	req := httptest.NewRequest("GET", "/service_members", nil)
	params := servicememberop.CreateServiceMemberParams{
		CreateServiceMemberPayload: &servicememberPayload,
		HTTPRequest:                req,
	}

	handler := CreateServiceMemberHandler(NewHandlerContext(suite.db, suite.logger))
	response := handler.Handle(params)

	_, ok := response.(*servicememberop.CreateServiceMemberUnauthorized)
	if !ok {
		t.Fatalf("Request failed: %#v", response)
	}
	// Then: we expect no servicemembers to have been created
	servicemembers := []models.ServiceMember{}
	suite.db.All(&servicemembers)

	if len(servicemembers) > 0 {
		t.Errorf("Expected to find no servicemembers but found %v", len(servicemembers))
	}
}

func (suite *HandlerSuite) TestPatchServiceMemberHandler() {
	t := suite.T()

	// Given: a logged in user
	user := models.User{
		LoginGovUUID:  uuid.Must(uuid.NewV4()),
		LoginGovEmail: "email@example.com",
		Type:          internalmessages.UserTypeUNKNOWN,
	}
	suite.mustSave(&user)

	// TODO: add more fields to change
	var origEdipi = "2342342344"
	var newEdipi = "9999999999"
	newServiceMember := models.ServiceMember{
		UserID: user.ID,
		Edipi:  &origEdipi,
	}
	suite.mustSave(&newServiceMember)

	patchPayload := internalmessages.PatchServiceMemberPayload{
		Edipi: &newEdipi,
	}

	// And: the context contains the auth values
	req := httptest.NewRequest("PATCH", "/service_members/some_id", nil)
	ctx := req.Context()
	ctx = context.PopulateAuthContext(ctx, user.ID, "fake token")
	req = req.WithContext(ctx)

	params := servicememberop.PatchServiceMemberParams{
		HTTPRequest:               req,
		ServiceMemberID:           strfmt.UUID(newServiceMember.ID.String()),
		PatchServiceMemberPayload: &patchPayload,
	}

	handler := PatchServiceMemberHandler(NewHandlerContext(suite.db, suite.logger))
	response := handler.Handle(params)

	okResponse, ok := response.(*servicememberop.PatchServiceMemberCreated)
	if !ok {
		t.Fatalf("Request failed: %#v", response)
	}

	patchServiceMemberPayload := okResponse.Payload

	if *patchServiceMemberPayload.Edipi != newEdipi {
		t.Fatalf("Edipi should have been updated.")
	}
}

func (suite *HandlerSuite) TestPatchServiceMemberHandlerWrongUser() {
	t := suite.T()

	// Given: a logged in user
	user := models.User{
		LoginGovUUID:  uuid.Must(uuid.NewV4()),
		LoginGovEmail: "email@example.com",
		Type:          internalmessages.UserTypeUNKNOWN,
	}
	suite.mustSave(&user)

	user2 := models.User{
		LoginGovUUID:  uuid.Must(uuid.NewV4()),
		LoginGovEmail: "email2@example.com",
		Type:          internalmessages.UserTypeUNKNOWN,
	}
	suite.mustSave(&user2)

	var origEdipi = "2342342344"
	var newEdipi = "9999999999"
	newServiceMember := models.ServiceMember{
		UserID: user.ID,
		Edipi:  &origEdipi,
	}
	suite.mustSave(&newServiceMember)

	patchPayload := internalmessages.PatchServiceMemberPayload{
		Edipi: &newEdipi,
	}

	// And: the context contains the auth values
	req := httptest.NewRequest("PATCH", "/service_members/some_id", nil)
	ctx := req.Context()
	ctx = context.PopulateAuthContext(ctx, user2.ID, "fake token")
	req = req.WithContext(ctx)

	params := servicememberop.PatchServiceMemberParams{
		HTTPRequest:               req,
		ServiceMemberID:           strfmt.UUID(newServiceMember.ID.String()),
		PatchServiceMemberPayload: &patchPayload,
	}

	handler := PatchServiceMemberHandler(NewHandlerContext(suite.db, suite.logger))
	response := handler.Handle(params)

	_, ok := response.(*servicememberop.PatchServiceMemberForbidden)
	if !ok {
		t.Fatalf("Request failed: %#v", response)
	}
}

func (suite *HandlerSuite) TestPatchServiceMemberHandlerNoServiceMember() {
	t := suite.T()

	// Given: a logged in user
	user := models.User{
		LoginGovUUID:  uuid.Must(uuid.NewV4()),
		LoginGovEmail: "email@example.com",
		Type:          internalmessages.UserTypeUNKNOWN,
	}
	suite.mustSave(&user)

	servicememberUUID := uuid.Must(uuid.NewV4())

	var newEdipi = "9999999999"

	patchPayload := internalmessages.PatchServiceMemberPayload{
		Edipi: &newEdipi,
	}

	// And: the context contains the auth values
	req := httptest.NewRequest("PATCH", "/service_members/some_id", nil)
	ctx := req.Context()
	ctx = context.PopulateAuthContext(ctx, user.ID, "fake token")
	req = req.WithContext(ctx)

	params := servicememberop.PatchServiceMemberParams{
		HTTPRequest:               req,
		ServiceMemberID:           strfmt.UUID(servicememberUUID.String()),
		PatchServiceMemberPayload: &patchPayload,
	}

	handler := PatchServiceMemberHandler(NewHandlerContext(suite.db, suite.logger))
	response := handler.Handle(params)

	_, ok := response.(*servicememberop.PatchServiceMemberNotFound)
	if !ok {
		t.Fatalf("Request failed: %#v", response)
	}
}

func (suite *HandlerSuite) TestPatchServiceMemberHandlerNoChange() {
	t := suite.T()

	// Given: a logged in user with a servicemember
	user := models.User{
		LoginGovUUID:  uuid.Must(uuid.NewV4()),
		LoginGovEmail: "email@example.com",
		Type:          internalmessages.UserTypeUNKNOWN,
	}
	suite.mustSave(&user)

	var origEdipi = "4444444444"
	newServiceMember := models.ServiceMember{
		UserID: user.ID,
		Edipi:  &origEdipi,
	}
	suite.mustSave(&newServiceMember)

	patchPayload := internalmessages.PatchServiceMemberPayload{}

	// And: the context contains the auth values
	req := httptest.NewRequest("PATCH", "/service_members/some_id", nil)
	ctx := req.Context()
	ctx = context.PopulateAuthContext(ctx, user.ID, "fake token")
	req = req.WithContext(ctx)

	params := servicememberop.PatchServiceMemberParams{
		HTTPRequest:               req,
		ServiceMemberID:           strfmt.UUID(newServiceMember.ID.String()),
		PatchServiceMemberPayload: &patchPayload,
	}

	handler := PatchServiceMemberHandler(NewHandlerContext(suite.db, suite.logger))
	response := handler.Handle(params)

	_, ok := response.(*servicememberop.PatchServiceMemberCreated)
	if !ok {
		t.Fatalf("Request failed: %#v", response)
	}
}
