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
      "title":"task ex1",
      "description":"desc demo"
   }

4. แก้ไขงาน
 `PUT /api/v1/tasks/<id>`
   ```JSON
   {
      "title":"task ex1 update",
      "description":"desc demo update",
      "status":"completed"
   }

5. ลบงาน
`DELETE /api/v1/tasks/<id>`


6. เข้าสู่ระบบ การเข้าสู่ระบบจะต้องใส่รหัส 2Fa Code ที่ได้รับจากเเอป Google Authentication twofactor สามารถนำเเอป Google Authentication twofactor ไปสเเกน QR Code เพื่อขอรหัสผ่านได้ที่ <http://localhost/auth/regis2Fa> โดยการกรอก username ที่ได้ลงทะเบียนไว้กับระบบเราไว้เเล้ว Task Management System
`POST /api/v1/signIn`
   ```JSON
   { 
      "username":"user1",
      "password":"user1@123",
      "otp_code": "971xxx"
   }


7. ลงทะเบียนผู้ใช้งานใหม่
`POST /api/v1/signUp`
   ```JSON
   {
   "full_name":"user1",
   "profile_image":"user1_profile_image.png",
   "username":"user1",
   "password":"password123xxx"
   }

<br/>

## เพิ่มเติม

1. การเข้าสู่ระบบ การเข้าสู่ระบบจะเพิ่มความปลอดภัยป้องกันการ Brute Force Attack Password  จึงกำหนดให้สามารถใส่รหัสผ่านผิดได้ไม่เกิน 5 ครั้ง ถ้าใส่รหัสผ่านผิดเกิน 5 ครั้ง ระบบจะ Lock Account นั้นไม่สามารถเข้าสู่ระบบได้ครั้งต่อไป เเต่ไม่รวมการใส่รหัส code ของ  2Fa code ที่ได้รับจาก Google Authentication twofactor ซึ่ง 2Fa code สามารถใส่ผิดได้หลายครั้ง เพราะรหัสค่อนข้าง reset เร็วจึงไม่ได้จำกัดในส่วนของ 2Fa code

2. การเข้าสู่ระบบโดยใช้ 2Fa  ให้ผู้ใช้งานดาวน์โหลด App Google Authentication twofactor จากนั้นลงทะเบียนเข้าใช้งานเเอปให้เรียบร้อย จากนั้นไปที่ปุ่มเพิ่มเพิ่ม Qr Code จากนั้นนำไปสเเกน Qr Code ที่หน้า <http://localhost/auth/regis2Fa> จะได้ 2fa code มาเพื่อนำไปเข้าสู่ระบบ 

<iframe src="https://github.com/nathakritbc/task-management/blob/main/pdf/auth2fa.pdf" width="640" height="480"></iframe>


  
