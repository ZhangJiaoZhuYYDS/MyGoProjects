package model

import (
    "time"
)

type Us struct{
    // 
    Id int64 `json:"id" form:"id" gorm:"primaryKey" `
    // 
    CreatedAt time.Time `json:"createdAt" form:"createdAt" gorm:"autoCreateTime" `
    // 
    UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" gorm:"autoUpdateTime" `
    // 
    DeletedAt gorm.DeletedAt `json:"deletedAt" form:"deletedAt" `
    
}
