# ระบบการจัดการรายการงาน (Task Management System)

## Stack ที่ใช้ในการพัฒนา
- `frontend`: Vue.js composition api
- `backend/`: Go Fiber Framework
- `database`: Postgresql Gorm orm
- `container` : Docker

## การติดตั้งและรันโครงการ

1. Clone Repository

   ```bash
   git clone https://github.com/nathakritbc/task-management.git
   ```

2. เข้าสู่ไดเรกทอรีของโปรเจ็ก

    ```bash
    cd task-management
    ```

3. รันคำสั่งต่อไปนี้ เพื่อ Build Docker:

   ```bash
   docker-compose up --build
   ```

4. ต่อมารันคำสั่งต่อไปนี้ เพื่อรัน Project:

    ```bash
   docker-compose up -d
   ```

5. ตรวจสอบ Container ของ Docker เเต่ล่ะตัวว่า Run สำเร็จหรือไม่:

6. จากนั้น ไปที่ <http://localhost> เพื่อเข้าใช้งานระบบ

## การเข้าถึงแอปพลิเคชัน

- Backend API: <http://localhost:9090>
- Frontend Application: <http://localhost>


## API Endpoints

1. แสดงรายการงานทั้งหมด
`GET /api/v1/tasks` 

2. ค้นหารายการงานเเละสถานะของงาน
`GET /api/v1/tasks?search=หัวข้องาน&status=in_process`

3. สร้างงานใหม่
 `POST /api/v1/tasks`
   ```JSON
   {
    "title":"Y1 top 23412",
    "description": "description1"
    }
```
 
4. แก้ไขงาน 
 `PUT /api/v1/tasks/<id>`


5. ลบงาน
`DELETE /api/v1/tasks/<id>/`

6. เข้าสู่ระบบ
`POST /api/v1/signIn`

7. ลงทะเบียนผู้ใช้งานใหม่
`POST /api/v1/signUp`: 

## ข้อสังเกต

- ตรวจสอบให้แน่ใจว่าได้ติดตั้ง lib modules ต่างๆครบถ้วยหรือไม่ก่อนรันโปรเจ็ก
- สำหรับการใช้งานในโหมดโปรดักชั่นให้เปลี่ยนรหัสผ่านและการตั้งค่าความปลอดภัยต่าง ๆ
