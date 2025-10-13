# Tổng Hợp Các Công Thức Tính Toán Trong Chương III

## Bảng Các Công Thức Chính

| **STT** | **Công Thức** | **Ý Nghĩa** | **Mục Đích Sử Dụng** |
|---------|---------------|-------------|---------------------|
| **1** | `LCOM2 = { P - Q, if P ≥ Q; 0, otherwise }` | **Lack of Cohesion in Methods** - Đo lường độ gắn kết của lớp | Định lượng độ gắn kết nội bộ của một module/lớp<br/>- **P**: Số cặp phương thức không chia sẻ thuộc tính<br/>- **Q**: Số cặp phương thức chia sẻ thuộc tính<br/>- **Giá trị cao** = độ gắn kết thấp |
| **2** | `I = Cₑ / (Cₑ + Cₐ)` | **Instability Index** - Chỉ số bất ổn định của component | Đo lường mức độ dễ thay đổi của một component<br/>- **Cₑ**: Efferent Coupling (phụ thuộc đi ra)<br/>- **Cₐ**: Afferent Coupling (phụ thuộc đi vào)<br/>- **I ≈ 0**: Rất ổn định (policy modules)<br/>- **I ≈ 1**: Rất bất ổn (detail modules) |

## Chi Tiết Các Công Thức

### 1. Công Thức LCOM2 (Lack of Cohesion in Methods)

```
LCOM2 = {
    P - Q,  if P ≥ Q
    0,      otherwise
}
```

**Trong đó:**
- **P** = Số cặp phương thức không chia sẻ thuộc tính (_attributes_)
- **Q** = Số cặp phương thức chia sẻ thuộc tính

**Giải thích:**
- **LCOM cao** → **Độ gắn kết thấp** → Cần refactor
- **LCOM thấp** → **Độ gắn kết cao** → Thiết kế tốt

### 2. Công Thức Instability Index (I)

```
I = Cₑ / (Cₑ + Cₐ)
```

**Trong đó:**
- **Cₑ** = Efferent Coupling (số kết nối đi ra)
- **Cₐ** = Afferent Coupling (số kết nối đi vào)

**Phân tích giá trị:**
- **I ≈ 0** (0 - 0.3): **Rất ổn định** - Policy modules, abstractions
- **I ≈ 1** (0.7 - 1): **Rất bất ổn** - Detail modules, implementations  
- **0.3 < I < 0.7**: **Trung bình** - Nên tránh

## Ứng Dụng Thực Tế

### Ngưỡng Quan Trọng

| **Chỉ Số** | **Ngưỡng Cảnh Báo** | **Hành Động Khuyến Nghị** |
|------------|---------------------|---------------------------|
| **Cₑ** | > 20 | Component được coi là **không ổn định** |
| **Cₐ** | Cao | Component **quan trọng** nhưng **cứng nhắc** |
| **I** | 0.3 < I < 0.7 | **Tránh** - Cần thiết kế lại |

### Liên Kết với Nguyên Tắc SOLID

| **Công Thức** | **Liên Quan SOLID** | **Ý Nghĩa Kiến Trúc** |
|---------------|---------------------|----------------------|
| **LCOM2** | **SRP** | Đảm bảo Functional Cohesion cao |
| **I ≈ 0** | **DIP** | Policy modules ổn định |
| **I ≈ 1** | **DIP** | Detail modules có thể thay đổi |
| **Cₑ thấp** | **OCP** | Bảo vệ core modules khỏi biến động |

## Ví Dụ Thực Tế Chi Tiết

### Ví Dụ 1: Tính LCOM2 cho Lớp Employee (Vi Phạm SRP)

**Tình huống:** Lớp `Employee` chứa 3 phương thức phục vụ 3 actor khác nhau:

```java
public class Employee {
    private String name;
    private double hourlyRate;
    private int hoursWorked;
    
    // Phương thức cho CFO
    public double calculatePay() { 
        return hourlyRate * hoursWorked; 
    }
    
    // Phương thức cho COO  
    public String reportHours() { 
        return "Hours worked: " + hoursWorked; 
    }
    
    // Phương thức cho CTO
    public void save() { 
        // Lưu vào database
    }
}
```

**Phân tích LCOM2:**
- **Phương thức:** `calculatePay()`, `reportHours()`, `save()`
- **Thuộc tính:** `name`, `hourlyRate`, `hoursWorked`

**Tính P (cặp phương thức KHÔNG chia sẻ thuộc tính):**
- `calculatePay()` ↔ `reportHours()`: Cả hai đều dùng `hoursWorked` → **Chia sẻ** ❌
- `calculatePay()` ↔ `save()`: `calculatePay()` dùng `hourlyRate, hoursWorked`; `save()` không dùng → **Không chia sẻ** ✅
- `reportHours()` ↔ `save()`: `reportHours()` dùng `hoursWorked`; `save()` không dùng → **Không chia sẻ** ✅

**P = 2**

**Tính Q (cặp phương thức CHIA SẺ thuộc tính):**
- `calculatePay()` ↔ `reportHours()`: Cả hai đều dùng `hoursWorked` → **Chia sẻ** ✅

**Q = 1**

**Kết quả LCOM2:**
```
LCOM2 = P - Q = 2 - 1 = 1
```

**Suy luận:** LCOM2 = 1 > 0 → **Độ gắn kết thấp** → Vi phạm SRP → Cần refactor

---

### Ví Dụ 2: Tính LCOM2 cho Lớp PayCalculator (Tuân Thủ SRP)

**Tình huống:** Sau khi refactor, tách thành lớp `PayCalculator`:

```java
public class PayCalculator {
    private double hourlyRate;
    private int hoursWorked;
    
    public PayCalculator(double hourlyRate, int hoursWorked) {
        this.hourlyRate = hourlyRate;
        this.hoursWorked = hoursWorked;
    }
    
    public double calculatePay() { 
        return hourlyRate * hoursWorked; 
    }
    
    public double calculateOvertime() {
        if (hoursWorked > 40) {
            return (hoursWorked - 40) * hourlyRate * 1.5;
        }
        return 0;
    }
}
```

**Phân tích LCOM2:**
- **Phương thức:** `calculatePay()`, `calculateOvertime()`
- **Thuộc tính:** `hourlyRate`, `hoursWorked`

**Tính P (cặp phương thức KHÔNG chia sẻ thuộc tính):**
- `calculatePay()` ↔ `calculateOvertime()`: Cả hai đều dùng `hourlyRate` và `hoursWorked` → **Chia sẻ** ❌

**P = 0**

**Tính Q (cặp phương thức CHIA SẺ thuộc tính):**
- `calculatePay()` ↔ `calculateOvertime()`: Cả hai đều dùng `hourlyRate` và `hoursWorked` → **Chia sẻ** ✅

**Q = 1**

**Kết quả LCOM2:**
```
LCOM2 = max(0, P - Q) = max(0, 0 - 1) = 0
```

**Suy luận:** LCOM2 = 0 → **Độ gắn kết cao** → Tuân thủ SRP → Thiết kế tốt

---

### Ví Dụ 3: Tính Instability Index cho Component DatabaseService

**Tình huống:** Hệ thống có các component sau:

```
DatabaseService (Core Policy)
├── UserRepository (depends on DatabaseService)
├── OrderRepository (depends on DatabaseService)  
├── ProductRepository (depends on DatabaseService)
└── MySQLConnection (depends on DatabaseService)

DatabaseService depends on:
├── ConnectionPool
└── Logger
```

**Phân tích Coupling cho DatabaseService:**

**Afferent Coupling (Cₐ) - Phụ thuộc đi vào:**
- `UserRepository` → `DatabaseService`
- `OrderRepository` → `DatabaseService`
- `ProductRepository` → `DatabaseService`
- `MySQLConnection` → `DatabaseService`

**Cₐ = 4**

**Efferent Coupling (Cₑ) - Phụ thuộc đi ra:**
- `DatabaseService` → `ConnectionPool`
- `DatabaseService` → `Logger`

**Cₑ = 2**

**Tính Instability Index:**
```
I = Cₑ / (Cₑ + Cₐ) = 2 / (2 + 4) = 2/6 = 0.33
```

**Suy luận:** 
- **I = 0.33** → Nằm trong khoảng **0.3 < I < 0.7**
- → **Cảnh báo:** Component có độ ổn định trung bình
- → **Khuyến nghị:** Cần thiết kế lại để đạt **I ≈ 0** (policy module) hoặc **I ≈ 1** (detail module)

---

### Ví Dụ 4: Tính Instability Index cho Component MySQLConnection (Detail Module)

**Tình huống:** Component `MySQLConnection`:

```
MySQLConnection depends on:
├── DatabaseService (interface)
├── MySQLDriver
├── ConnectionString
└── TimeoutConfig

Components depending on MySQLConnection:
├── DatabaseService (implementation)
```

**Phân tích Coupling cho MySQLConnection:**

**Afferent Coupling (Cₐ):**
- `DatabaseService` → `MySQLConnection`

**Cₐ = 1**

**Efferent Coupling (Cₑ):**
- `MySQLConnection` → `DatabaseService`
- `MySQLConnection` → `MySQLDriver`
- `MySQLConnection` → `ConnectionString`
- `MySQLConnection` → `TimeoutConfig`

**Cₑ = 4**

**Tính Instability Index:**
```
I = Cₑ / (Cₑ + Cₐ) = 4 / (4 + 1) = 4/5 = 0.8
```

**Suy luận:**
- **I = 0.8** → Nằm trong khoảng **0.7 ≤ I ≤ 1**
- → **Tốt:** Component có độ bất ổn cao (detail module)
- → **Phù hợp:** Implementation modules nên có I cao để dễ thay đổi

---

### Ví Dụ 5: So Sánh Trước và Sau Refactoring

**Trước Refactoring - Lớp UserService (Vi phạm SRP):**

```java
public class UserService {
    private UserRepository userRepo;
    private EmailService emailService;
    private Logger logger;
    private DatabaseConnection db;
    
    // Nhiều phương thức không liên quan
    public void createUser() { /* uses userRepo, db */ }
    public void sendWelcomeEmail() { /* uses emailService */ }
    public void logActivity() { /* uses logger */ }
    public void backupDatabase() { /* uses db */ }
}
```

**LCOM2 Analysis:**
- **Phương thức:** 4 methods
- **Thuộc tính:** 4 fields
- **P = 4** (nhiều cặp không chia sẻ thuộc tính)
- **Q = 2** (ít cặp chia sẻ thuộc tính)
- **LCOM2 = 4 - 2 = 2** → **Độ gắn kết thấp**

**Sau Refactoring - Tách thành các lớp riêng:**

```java
// UserManagement (High Cohesion)
public class UserManagement {
    private UserRepository userRepo;
    private DatabaseConnection db;
    
    public void createUser() { /* uses userRepo, db */ }
    public void updateUser() { /* uses userRepo, db */ }
}

// EmailNotification (High Cohesion)  
public class EmailNotification {
    private EmailService emailService;
    
    public void sendWelcomeEmail() { /* uses emailService */ }
    public void sendPasswordReset() { /* uses emailService */ }
}
```

**LCOM2 Analysis cho UserManagement:**
- **Phương thức:** 2 methods
- **Thuộc tính:** 2 fields
- **P = 0** (cả hai đều dùng cùng thuộc tính)
- **Q = 1** (cả hai chia sẻ thuộc tính)
- **LCOM2 = 0** → **Độ gắn kết cao**

---

## Kết Luận Thực Tế

### Khi Nào Sử Dụng Các Công Thức:

1. **LCOM2**: Sử dụng khi **code review** hoặc **refactoring** để đánh giá độ gắn kết của lớp
2. **Instability Index**: Sử dụng khi **thiết kế kiến trúc** để đảm bảo phân tách đúng giữa policy và detail modules

### Quy Trình Thực Tế:

1. **Đo lường** → **Phân tích** → **Suy luận** → **Hành động**
2. **LCOM2 cao** → Refactor để tách lớp
3. **I trung bình** → Thiết kế lại để đạt phân cực (I≈0 hoặc I≈1)
4. **Tích hợp vào CI/CD** để theo dõi chất lượng liên tục

## Tóm Tắt

Các công thức này tạo thành một **hệ thống đo lường định lượng** cho chất lượng kiến trúc, cho phép kiến trúc sư:

1. **Đánh giá khách quan** chất lượng thiết kế
2. **Xác định vấn đề** cần refactor
3. **Đảm bảo tuân thủ** các nguyên tắc SOLID
4. **Duy trì tính ổn định** của kiến trúc theo thời gian
