package dao

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"gocument/app/api/global"
	"gocument/app/api/internal/model"
	"golang.org/x/net/context"
)

func InsertDocument(document *model.Document, documentContent *model.DocumentContent) error {
	result := global.MysqlDB.Model(model.Document{}).Create(document)
	if result.Error != nil {
		return result.Error
	}
	documentContent.DocID = document.ID
	database := global.MongoDB.Database("gocument")
	collection := database.Collection("documents")
	_, err := collection.InsertOne(context.TODO(), *documentContent)
	if err != nil {
		return err
	}
	return nil
}

func FindDocument(document *model.Document, wholeDocument *model.WholeDocument) error {
	result := global.MysqlDB.Model(model.Document{}).Where("id = ? AND user_id = ?", document.ID, document.UserID).First(document)
	// if result.RowsAffected == 0 {
	// 	return fmt.Errorf("文档不存在")
	// } else
	if result.Error != nil {
		return result.Error
	}
	collection := global.MongoDB.Database("gocument").Collection("documents")
	var documentContent model.DocumentContent
	err := collection.FindOne(context.TODO(), bson.M{"doc_id": document.ID}).Decode(&documentContent)
	if err != nil {
		return err
	}
	wholeDocument.Document = *document
	wholeDocument.Content = documentContent
	return nil
}
func FindDocumentByID(document *model.Document, wholeDocument *model.WholeDocument) error {
	result := global.MysqlDB.Model(model.Document{}).Where("id = ?", document.ID).First(document)
	if result.RowsAffected == 0 {
		return fmt.Errorf("文档不存在")
	} else if result.Error != nil {
		return result.Error
	}
	collection := global.MongoDB.Database("gocument").Collection("documents")
	var documentContent model.DocumentContent
	err := collection.FindOne(context.TODO(), bson.M{"doc_id": document.ID}).Decode(&documentContent)
	if err != nil {
		return err
	}
	wholeDocument.Document = *document
	wholeDocument.Content = documentContent
	return nil
}
func UpdateDocument(document *model.Document, documentContent *model.DocumentContent) error {
	result := global.MysqlDB.Model(model.Document{}).Where("id = ? AND user_id = ?", document.ID, document.UserID).Updates(document)
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到你的文档")
	}
	if result.Error != nil {
		return result.Error
	}
	collection := global.MongoDB.Database("gocument").Collection("documents")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"doc_id": document.ID}, bson.M{"$set": bson.M{"content": documentContent.Content}})
	if err != nil {
		return err
	}
	return nil
}

func GetDocumentList(document model.Document, documentList *[]model.Document) error {
	result := global.MysqlDB.Model(model.Document{}).Where("user_id = ?", document.UserID).Find(documentList)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteDocument(document *model.Document) error {
	result := global.MysqlDB.Model(model.Document{}).Where("id = ? AND user_id = ?", document.ID, document.UserID).Delete(document)
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到你的文档")
	}
	if result.Error != nil {
		return result.Error
	}
	collection := global.MongoDB.Database("gocument").Collection("documents")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"doc_id": document.ID})
	if err != nil {
		return err
	}
	return nil
}
