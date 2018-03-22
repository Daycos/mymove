package handlers

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/satori/go.uuid"

	uploadop "github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/uploads"
	// "github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

type FakeS3 struct {
	putFiles []*s3.PutObjectInput
}

func (fake *FakeS3) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	fake.putFiles = append(fake.putFiles, input)
	s3PutObjectOutput := &s3.PutObjectOutput{}
	return s3PutObjectOutput, nil
}

func (suite *HandlerSuite) TestCreateUploadsHandler() {
	t := suite.T()

	move, err := testdatagen.MakeMove(suite.db)
	if err != nil {
		t.Fatalf("could not create move: %s", err)
	}

	document, err := testdatagen.MakeDocument(suite.db, &move)
	if err != nil {
		t.Fatalf("could not create document: %s", err)
	}

	fakeS3 := &FakeS3{}
	params := uploadop.NewCreateUploadParams()
	params.MoveID = *fmtUUID(move.ID)
	params.DocumentID = *fmtUUID(document.ID)
	handler := CreateUploadHandler(NewS3HandlerContext(suite.db, suite.logger, fakeS3))
	response := handler.Handle(params)

	createdResponse, ok := response.(*uploadop.CreateUploadCreated)
	if !ok {
		t.Fatalf("Request failed: %#v", response)
	}
	uploadPayload := createdResponse.Payload

	if uuid.Must(uuid.FromString(uploadPayload.ID.String())) == uuid.Nil {
		t.Errorf("got empty document uuid")
	}
}
