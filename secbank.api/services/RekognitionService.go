package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/google/uuid"
	"strconv"
)

type RekognitionService struct {
	Client *rekognition.Rekognition
}

// Função para inicializar o RekognitionService
func NewRekognitionService(region string) *RekognitionService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewEnvCredentials(),
	}))

	return &RekognitionService{
		Client: rekognition.New(sess),
	}
}

// Método para criar uma coleção
func (r *RekognitionService) CreateCollection() (string, error) {
	collectionID := uuid.New().String()
	_, err := r.Client.CreateCollection(&rekognition.CreateCollectionInput{
		CollectionId: aws.String(collectionID),
	})
	if err != nil {
		return "", err
	}
	return collectionID, nil
}

// Método para criar um usuário
func (r *RekognitionService) CreateUser(collectionID string, customerID int) error {
	_, err := r.Client.CreateUser(&rekognition.CreateUserInput{
		CollectionId: aws.String(collectionID),
		UserId:       aws.String(strconv.Itoa(customerID)),
	})
	if err != nil {
		return err
	}
	return nil
}

// Método para associar faces a um usuário
func (r *RekognitionService) AssociateFacesToUser(collectionID string, userID int, faceIDs []string) error {
	_, err := r.Client.AssociateFaces(&rekognition.AssociateFacesInput{
		CollectionId: aws.String(collectionID),
		UserId:       aws.String(strconv.Itoa(userID)),
		FaceIds:      aws.StringSlice(faceIDs),
	})
	return err
}

// Método para indexar faces a partir de uma imagem codificada em base64
func (r *RekognitionService) IndexFaces(collectionID string, imageBytes []byte) (*rekognition.IndexFacesOutput, error) {
	image := &rekognition.Image{
		Bytes: imageBytes,
	}
	response, err := r.Client.IndexFaces(&rekognition.IndexFacesInput{
		CollectionId: aws.String(collectionID),
		Image:        image,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Método para buscar usuários por uma imagem codificada em base64
func (r *RekognitionService) SearchUsersByImage(collectionID string, imageBytes []byte) (*rekognition.SearchUsersByImageOutput, error) {
	image := &rekognition.Image{
		Bytes: imageBytes,
	}
	response, err := r.Client.SearchUsersByImage(&rekognition.SearchUsersByImageInput{
		CollectionId: aws.String(collectionID),
		Image:        image,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *RekognitionService) GetFacesIDs(responseToIndexFace *rekognition.IndexFacesOutput) []string {
	var faceIDs []string

	for _, faceRecord := range responseToIndexFace.FaceRecords {
		if faceRecord.Face != nil && faceRecord.Face.FaceId != nil {
			faceIDs = append(faceIDs, *faceRecord.Face.FaceId)
		}
	}

	return faceIDs
}
