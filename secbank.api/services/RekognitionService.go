package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/rekognition/rekognitioniface"
	"github.com/google/uuid"
	"strconv"
)

// RekognitionService encapsula operações do Rekognition
type RekognitionService struct {
	Client rekognitioniface.RekognitionAPI
}

// Inicializa RekognitionService com uma sessão AWS
func NewRekognitionService(region string) *RekognitionService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewEnvCredentials(),
	}))
	return &RekognitionService{Client: rekognition.New(sess)}
}

// Cria uma coleção e retorna seu ID
func (r *RekognitionService) CreateCollection() (string, error) {
	collectionID := uuid.New().String()
	_, err := r.Client.CreateCollection(&rekognition.CreateCollectionInput{
		CollectionId: aws.String(collectionID),
	})
	return collectionID, err
}

// Cria um usuário em uma coleção
func (r *RekognitionService) CreateUser(collectionID string, customerID int) error {
	_, err := r.Client.CreateUser(&rekognition.CreateUserInput{
		CollectionId: aws.String(collectionID),
		UserId:       aws.String(strconv.Itoa(customerID)),
	})
	return err
}

// Associa faces a um usuário
func (r *RekognitionService) AssociateFacesToUser(collectionID string, userID int, faceIDs []string) error {
	_, err := r.Client.AssociateFaces(&rekognition.AssociateFacesInput{
		CollectionId: aws.String(collectionID),
		UserId:       aws.String(strconv.Itoa(userID)),
		FaceIds:      aws.StringSlice(faceIDs),
	})
	return err
}

// Indexa faces em uma coleção
func (r *RekognitionService) IndexFaces(collectionID string, imageBytes []byte) (*rekognition.IndexFacesOutput, error) {
	return r.Client.IndexFaces(&rekognition.IndexFacesInput{
		CollectionId: aws.String(collectionID),
		Image:        &rekognition.Image{Bytes: imageBytes},
	})
}

// Busca usuários por uma imagem
func (r *RekognitionService) SearchUsersByImage(collectionID string, imageBytes []byte) (*rekognition.SearchUsersByImageOutput, error) {
	return r.Client.SearchUsersByImage(&rekognition.SearchUsersByImageInput{
		CollectionId: aws.String(collectionID),
		Image:        &rekognition.Image{Bytes: imageBytes},
	})
}

// Obtém os IDs das faces indexadas
func (r *RekognitionService) GetFacesIDs(indexOutput *rekognition.IndexFacesOutput) []string {
	faceIDs := make([]string, len(indexOutput.FaceRecords))
	for i, faceRecord := range indexOutput.FaceRecords {
		if faceRecord.Face != nil && faceRecord.Face.FaceId != nil {
			faceIDs[i] = *faceRecord.Face.FaceId
		}
	}
	return faceIDs
}
