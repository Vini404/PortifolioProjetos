package test

import (
	"secbank.api/services"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/rekognition/rekognitioniface"
	"github.com/stretchr/testify/assert"
)

// Mock para simular o comportamento do RekognitionAPI
type mockRekognitionClient struct {
	rekognitioniface.RekognitionAPI
	CreateCollectionOutput   *rekognition.CreateCollectionOutput
	CreateCollectionError    error
	CreateUserOutput         *rekognition.CreateUserOutput
	CreateUserError          error
	AssociateFacesOutput     *rekognition.AssociateFacesOutput
	AssociateFacesError      error
	IndexFacesOutput         *rekognition.IndexFacesOutput
	IndexFacesError          error
	SearchUsersByImageOutput *rekognition.SearchUsersByImageOutput
	SearchUsersByImageError  error
}

func (m *mockRekognitionClient) CreateCollection(input *rekognition.CreateCollectionInput) (*rekognition.CreateCollectionOutput, error) {
	return m.CreateCollectionOutput, m.CreateCollectionError
}

func (m *mockRekognitionClient) CreateUser(input *rekognition.CreateUserInput) (*rekognition.CreateUserOutput, error) {
	return m.CreateUserOutput, m.CreateUserError
}

func (m *mockRekognitionClient) AssociateFaces(input *rekognition.AssociateFacesInput) (*rekognition.AssociateFacesOutput, error) {
	return m.AssociateFacesOutput, m.AssociateFacesError
}

func (m *mockRekognitionClient) IndexFaces(input *rekognition.IndexFacesInput) (*rekognition.IndexFacesOutput, error) {
	return m.IndexFacesOutput, m.IndexFacesError
}

func (m *mockRekognitionClient) SearchUsersByImage(input *rekognition.SearchUsersByImageInput) (*rekognition.SearchUsersByImageOutput, error) {
	return m.SearchUsersByImageOutput, m.SearchUsersByImageError
}

func TestCreateCollection(t *testing.T) {
	mockClient := &mockRekognitionClient{
		CreateCollectionOutput: &rekognition.CreateCollectionOutput{},
		CreateCollectionError:  nil,
	}

	service := &services.RekognitionService{Client: mockClient}

	collectionID, err := service.CreateCollection()
	assert.NoError(t, err)
	assert.NotEmpty(t, collectionID, "Collection ID should not be empty")
}

func TestCreateUser(t *testing.T) {
	mockClient := &mockRekognitionClient{
		CreateUserOutput: &rekognition.CreateUserOutput{},
		CreateUserError:  nil,
	}

	service := &services.RekognitionService{Client: mockClient}

	err := service.CreateUser("test-collection", 123)
	assert.NoError(t, err)
}

func TestAssociateFacesToUser(t *testing.T) {
	mockClient := &mockRekognitionClient{
		AssociateFacesOutput: &rekognition.AssociateFacesOutput{},
		AssociateFacesError:  nil,
	}

	service := &services.RekognitionService{Client: mockClient}

	err := service.AssociateFacesToUser("test-collection", 123, []string{"face1", "face2"})
	assert.NoError(t, err)
}

func TestIndexFaces(t *testing.T) {
	mockClient := &mockRekognitionClient{
		IndexFacesOutput: &rekognition.IndexFacesOutput{
			FaceRecords: []*rekognition.FaceRecord{
				{
					Face: &rekognition.Face{
						FaceId: aws.String("face-id-123"),
					},
				},
			},
		},
		IndexFacesError: nil,
	}

	service := &services.RekognitionService{Client: mockClient}

	imageData := []byte("test-image-data")
	response, err := service.IndexFaces("test-collection", imageData)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Len(t, response.FaceRecords, 1)
}

func TestSearchUsersByImage(t *testing.T) {
	mockClient := &mockRekognitionClient{
		SearchUsersByImageOutput: &rekognition.SearchUsersByImageOutput{},
		SearchUsersByImageError:  nil,
	}

	service := &services.RekognitionService{Client: mockClient}

	imageData := []byte("test-image-data")
	response, err := service.SearchUsersByImage("test-collection", imageData)
	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestGetFacesIDs(t *testing.T) {
	mockClient := &mockRekognitionClient{
		IndexFacesOutput: &rekognition.IndexFacesOutput{
			FaceRecords: []*rekognition.FaceRecord{
				{
					Face: &rekognition.Face{
						FaceId: aws.String("face-id-123"),
					},
				},
			},
		},
	}

	service := &services.RekognitionService{Client: mockClient}

	response := &rekognition.IndexFacesOutput{
		FaceRecords: []*rekognition.FaceRecord{
			{
				Face: &rekognition.Face{
					FaceId: aws.String("face-id-123"),
				},
			},
		},
	}

	faceIDs := service.GetFacesIDs(response)
	assert.Len(t, faceIDs, 1)
	assert.Equal(t, "face-id-123", faceIDs[0])
}
