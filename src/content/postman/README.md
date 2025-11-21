# Course API Postman Collection

## 📋 Tổng quan

Postman Collection này chứa đầy đủ các API endpoints để test Course CRUD operations với search và pagination.

## 🚀 Cách sử dụng

### 1. Import Collection vào Postman

1. Mở Postman
2. Click **Import** button
3. Chọn file `Course_API_Collection.json`
4. Collection sẽ được import với tên "Course API Collection"

### 2. Cấu hình Environment

Collection sử dụng các variables:
- `baseUrl`: `http://localhost:9100` (mặc định)
- `courseId`: Tự động set từ response của Create Course

### 3. Thứ tự test được khuyến nghị

#### Phase 1: Setup Sample Data
1. **Create Sample Courses** (3 requests) - Tạo dữ liệu mẫu
2. **List All Courses** - Kiểm tra dữ liệu đã tạo

#### Phase 2: CRUD Operations
3. **Create Course** - Tạo course mới
4. **Get Course Detail** - Lấy chi tiết course (sử dụng courseId từ bước 3)
5. **Update Course** - Cập nhật course
6. **List Courses with Pagination** - Test pagination

#### Phase 3: Search & Filter
7. **Search Courses by Title** - Tìm kiếm theo title
8. **Search Courses by Instructor** - Lọc theo instructor
9. **Search Courses by Structure Type** - Lọc theo structure type
10. **Complex Search** - Tìm kiếm phức tạp

#### Phase 4: Delete Operations
11. **Delete Single Course** - Xóa 1 course
12. **Delete Multiple Courses** - Xóa nhiều courses

#### Phase 5: Error Handling
13. **Error Cases & Validation** - Test các trường hợp lỗi

## 📊 API Endpoints Overview

| Method | Endpoint | Description | Parameters |
|--------|----------|-------------|------------|
| POST | `/api/courses` | Create course | Body: Course data (snake_case) |
| GET | `/api/courses/{id}` | Get course detail | Path: courseId |
| GET | `/api/courses` | List courses | Query: title, instructor_id, structure_type, page, size |
| PUT | `/api/courses/{id}` | Update course | Path: courseId, Body: Update data (snake_case) |
| DELETE | `/api/courses` | Delete courses | Body: Array of courseIds |

## 🔍 Search Parameters

### GET /api/courses
- `title` (optional): Search by title (partial match, case-insensitive)
- `instructor_id` (optional): Filter by instructor UUID
- `structure_type` (optional): Filter by structure type (LINEAR/ADAPTIVE)
- `page` (default: 0): Page number (0-based)
- `size` (default: 20): Page size (max: 100)

### Examples:
```
GET /api/courses?title=java&page=0&size=10
GET /api/courses?instructor_id=660e8400-e29b-41d4-a716-446655440001
GET /api/courses?structure_type=LINEAR&page=0&size=5
```

## 📝 Request/Response Format

### Create Course Request:
```json
{
  "title": "Java Programming",
  "description": "Learn Java programming",
  "instructor_id": "660e8400-e29b-41d4-a716-446655440001",
  "structure_type": "LINEAR"
}
```

### Update Course Request:
```json
{
  "title": "Updated Title",
  "description": "Updated description",
  "structure_type": "ADAPTIVE"
}
```

### Delete Courses Request:
```json
[
  "550e8400-e29b-41d4-a716-446655440001",
  "550e8400-e29b-41d4-a716-446655440002"
]
```

### Standard Response Format:
```json
{
  "error_code": 0,
  "message": "Success",
  "data": {
    "course_id": "550e8400-e29b-41d4-a716-446655440001",
    "title": "Java Programming",
    "description": "Learn Java programming",
    "instructor_id": "660e8400-e29b-41d4-a716-446655440001",
    "structure_type": "LINEAR",
    "created_at": "2024-01-15T10:30:00",
    "updated_at": "2024-01-15T10:30:00"
  }
}
```

### Paginated Response Format:
```json
{
  "error_code": 0,
  "message": "Success",
  "data": {
    "content": [...],
    "page": 0,
    "size": 20,
    "total_elements": 50,
    "total_pages": 3
  }
}
```

## ⚠️ Error Cases

### Validation Errors (400):
- Missing required fields (title, description, instructor_id, structure_type)
- Invalid structure type (must be LINEAR or ADAPTIVE)
- Empty delete list

### Not Found Errors (404):
- Course ID not found in database
- Invalid UUID format

## 🧪 Automated Tests

Collection bao gồm automated tests:
- Status code validation
- Response format validation
- Auto-extract courseId từ create response
- Console logging cho debugging

## 🔧 Prerequisites

1. **Application running**: Ensure Content Service is running on port 9100
2. **Database**: PostgreSQL with courses table created
3. **Sample data**: Run SQL script để có dữ liệu mẫu

```bash
# Start application
mvn spring-boot:run

# Create database table
psql -U postgres -d co3017 -f src/main/resources/sql/create_courses_table.sql
```

## 📚 Tips & Best Practices

1. **Run in order**: Follow the recommended test sequence
2. **Check variables**: Monitor `courseId` variable after create operations
3. **Review responses**: Check both success and error responses
4. **Test edge cases**: Use error test cases để validate error handling
5. **Performance**: Test với large datasets để validate pagination

## 🐛 Troubleshooting

### Common Issues:
1. **Connection refused**: Check if application is running on port 9100
2. **404 errors**: Verify courseId exists in database
3. **Validation errors**: Check request body format và required fields
4. **Database errors**: Ensure PostgreSQL is running và courses table exists

### Debug Steps:
1. Check application logs
2. Verify database connection
3. Test with sample data first
4. Use Postman Console để xem detailed requests/responses
