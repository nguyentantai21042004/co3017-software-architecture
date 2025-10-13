# Chương 2 - Các Nguyên tắc SOLID trong Kiến Trúc Phần Mềm

## 1. Khái quát về SOLID và Mục tiêu Thiết kế Hệ thống

### 1.1. Bối cảnh: Từ Mã Sạch đến Kiến trúc Bền vững

Các hệ thống phần mềm chất lượng cao luôn bắt nguồn từ mã nguồn sạch và dễ bảo trì. Tuy nhiên, việc sở hữu các thành phần mã nguồn tốt (ví như "những viên gạch tốt") không tự động đảm bảo một kiến trúc tổng thể tốt. Lập trình viên có thể tạo ra một mớ hỗn độn đáng kể ngay cả khi sử dụng các khối mã nguồn được xây dựng tốt. 

Đây chính là lý do các **Nguyên tắc SOLID** được đưa ra, đóng vai trò như một bộ khung hướng dẫn thiết kế ở cấp độ trung gian (_mid-level software structures_).

**SOLID** là một bộ năm nguyên tắc thiết kế hướng đối tượng được đề xuất bởi **Robert C. Martin**. Các nguyên tắc này quy định cách thức sắp xếp các hàm và cấu trúc dữ liệu vào trong các lớp, và quan trọng hơn là cách các lớp đó nên được kết nối với nhau trong hệ thống.

### 1.2. Mục tiêu Kiến trúc Cấp trung của SOLID

Mục tiêu chính của việc áp dụng SOLID không chỉ dừng lại ở việc cải thiện chất lượng mã nguồn cục bộ, mà còn hướng đến việc xây dựng các cấu trúc phần mềm có khả năng đối phó với sự thay đổi của yêu cầu nghiệp vụ. 

Cụ thể, các nguyên tắc này hướng tới việc tạo ra các cấu trúc đạt được **ba thuộc tính cốt lõi**:

- **Dung thứ Thay đổi** (_Tolerate Change_): Hệ thống có khả năng thích nghi với các thay đổi yêu cầu mà không phải tái cấu trúc toàn bộ.

- **Dễ hiểu** (_Are Easy to Understand_): Cấu trúc mã nguồn rõ ràng, giúp nhà phát triển mới dễ dàng tiếp cận và bảo trì.

- **Là nền tảng của các thành phần tái sử dụng** (_Reusable Components_): Các lớp và module được thiết kế theo SOLID có thể được đóng gói và tái sử dụng hiệu quả trong nhiều hệ thống khác nhau.

Việc tuân thủ SOLID giúp giảm thiểu chi phí vòng đời của hệ thống và tối đa hóa năng suất của lập trình viên.   

---

## 2. Nguyên tắc Trách nhiệm Đơn lẻ (Single Responsibility Principle - SRP)

### 2.1. Định nghĩa Cốt lõi và Khái niệm Actor

**Nguyên tắc Trách nhiệm Đơn lẻ (SRP)** tuyên bố rằng mỗi module phần mềm chỉ nên có một, và chỉ một, lý do để thay đổi. 

Trong ngữ cảnh thiết kế hiện đại, định nghĩa này được diễn giải chặt chẽ hơn thông qua khái niệm **Actor**: Một module chỉ nên chịu trách nhiệm với một, và chỉ một, actor (tác nhân).

> **Actor (tác nhân)** được hiểu là người dùng, vai trò, hoặc nhóm chức năng có quyền yêu cầu hoặc chịu trách nhiệm về các thay đổi đối với chức năng của module đó.

Bằng cách giới hạn trách nhiệm chỉ với một actor, SRP đảm bảo rằng sự thay đổi được yêu cầu bởi một nhóm sẽ không gây ra ảnh hưởng bất ngờ hoặc không mong muốn đối với các nhóm khác.

### 2.2. Mối liên hệ với Cấu trúc Tổ chức (Định luật Conway)

SRP được xem là một hệ quả tích cực của **Định luật Conway**, phát biểu rằng:

> Cấu trúc tốt nhất cho một hệ thống phần mềm bị ảnh hưởng nặng nề bởi cấu trúc xã hội của tổ chức sử dụng nó.

Khi phân tách các trách nhiệm của mã nguồn theo các actor khác nhau (ví dụ: Tách logic Tài chính khỏi logic Vận hành và khỏi logic Cơ sở dữ liệu), thực chất là đang ánh xạ cấu trúc tổ chức vào cấu trúc mã nguồn. 

Điều này trực tiếp giảm thiểu xung đột giữa các nhóm phát triển khác nhau và tăng cường sự độc lập của các module.

### 2.3. Phân tích Trường hợp Vi phạm Điển hình (Anti-pattern)

#### Ví dụ: Lớp Employee vi phạm SRP

Xét lớp `Employee` vi phạm SRP, chứa ba phương thức phục vụ cho ba actor hoàn toàn khác nhau:

```java
public class Employee {
    // Cho CFO - Giám đốc Tài chính
    public Money calculatePay() { ... }
    
    // Cho COO - Giám đốc Vận hành  
    public String reportHours() { ... }
    
    // Cho CTO - Giám đốc Công nghệ
    public void save() { ... }
}
```

#### Hậu quả của Sự Vi phạm

Sự vi phạm này tạo ra **sự kết nối ngầm** (_incidental coupling_) giữa các chức năng độc lập. 

**Ví dụ cụ thể:** Giả sử hàm `calculatePay()` và `reportHours()` cùng chia sẻ một thuật toán chung để tính giờ làm việc không phải giờ tăng ca.

Nếu nhóm CFO yêu cầu điều chỉnh cách tính giờ không tăng ca trong `calculatePay()`, nhà phát triển thực hiện thay đổi đó có thể không biết rằng nhóm COO cũng đang sử dụng chính thuật toán này thông qua `reportHours()` cho mục đích khác. 

**Kết quả:** Chức năng của COO bị sửa đổi mà họ không hay biết, dẫn đến lỗi hoặc hành vi sai lệch trong môi trường vận hành của COO.

> **Nguyên tắc cốt lõi:** Nếu một module chịu trách nhiệm trước nhiều actor, mọi thay đổi từ một actor đều có nguy cơ lan truyền và gây ra lỗi cho các actor khác, làm giảm khả năng bảo trì và tính ổn định của hệ thống.

### 2.4. Các Giải pháp Thiết kế để Đảm bảo SRP

Để cô lập các lý do thay đổi, kiến trúc sư cần tách biệt các trách nhiệm.

#### Giải pháp 1: Tách biệt Chức năng (Refactoring)

Phương pháp cơ bản là tái cấu trúc ba phương thức thành ba lớp riêng biệt:

```java
// Trách nhiệm với CFO
public class PayCalculator {
    public Money calculatePay(EmployeeData data) { ... }
}

// Trách nhiệm với COO  
public class HourReporter {
    public String reportHours(EmployeeData data) { ... }
}

// Trách nhiệm với CTO
public class EmployeeSaver {
    public void save(EmployeeData data) { ... }
}
```

Ba lớp này hoàn toàn không biết về nhau và truy cập độc lập vào nguồn dữ liệu nhân viên chung (`EmployeeData`) để thực hiện hành vi của mình.

#### Giải pháp 2: Sử dụng Mẫu Thiết kế Facade (Mặt tiền)

Mặc dù Giải pháp 1 đảm bảo SRP, nó lại làm tăng độ phức tạp cho Client. Client giờ đây phải khởi tạo và theo dõi ba lớp riêng biệt để thực hiện các thao tác trên nhân viên.   

Để giải quyết vấn đề quản lý phức tạp này, **Mẫu Thiết kế Facade** được áp dụng:

```java
public class EmployeeFacade {
    private PayCalculator payCalculator;
    private HourReporter hourReporter;
    private EmployeeSaver employeeSaver;
    
    public Money calculatePay(EmployeeData data) {
        return payCalculator.calculatePay(data);
    }
    
    public String reportHours(EmployeeData data) {
        return hourReporter.reportHours(data);
    }
    
    public void save(EmployeeData data) {
        employeeSaver.save(data);
    }
}
```

Lớp `EmployeeFacade` được tạo ra, cung cấp một giao diện đơn giản hóa cho Client. Lớp Facade này che giấu sự phức tạp của hệ thống con (_Subsystems_) bằng cách ủy quyền các lệnh gọi đến các lớp thành phần đã được tách biệt.

### 2.5. Sự Mở rộng của SRP lên Cấp độ Kiến trúc

Sự ảnh hưởng của SRP không chỉ giới hạn ở cấp độ lớp. Nó tái xuất hiện ở các cấp độ kiến trúc cao hơn:

#### Cấp độ Thành phần (Component Level)
SRP trở thành **Nguyên tắc Bao đóng Chung** (_Common Closure Principle - CCP_). 

> **CCP** quy định rằng các lớp thay đổi cùng nhau (vì cùng một actor) nên được đóng gói cùng nhau trong cùng một thành phần.

#### Cấp độ Kiến trúc (Architectural Level)  
SRP trở thành **Trục Thay đổi** (_Axis of Change_), là động lực chính để xác định và thiết lập **Ranh giới Kiến trúc** (_Architectural Boundaries_).

Việc này đảm bảo rằng các module chứa logic cốt lõi ổn định được bảo vệ khỏi các chi tiết triển khai dễ thay đổi.

---

## Tóm tắt

Nguyên tắc SRP là nền tảng cho việc thiết kế hệ thống phần mềm bền vững. Bằng cách đảm bảo mỗi module chỉ phục vụ một actor duy nhất, chúng ta tạo ra các hệ thống dễ bảo trì, ít lỗi và có khả năng thích nghi với sự thay đổi.

## 3. Nguyên tắc Mở/Đóng (Open-Closed Principle - OCP)

### 3.1. Định nghĩa Cốt lõi

**Nguyên tắc Mở/Đóng (OCP)**, được đặt ra bởi **Bertrand Meyer** vào năm 1988, là nguyên tắc thứ hai trong bộ SOLID. 

> **OCP** tuyên bố rằng một thực thể phần mềm (lớp, module, hàm) nên **mở để mở rộng** (_Open for extension_) nhưng **đóng để sửa đổi** (_Closed for modification_).

**Mục tiêu cơ bản** là cho phép thêm các chức năng mới vào một thành phần mà không cần thay đổi mã nguồn hiện có của nó. Điều này là tối quan trọng để ngăn chặn các thay đổi lan truyền (_ripple out_) đến các lớp, đối tượng, hoặc module phụ thuộc.

### 3.2. Phân tích Tác động của Vi phạm OCP: Trường hợp Báo cáo Tài chính

#### Tình huống ban đầu:
Xét một hệ thống ban đầu chỉ phục vụ hiển thị tóm tắt tài chính trên một trang web:
- Dữ liệu cuộn được
- Số âm hiển thị màu đỏ

#### Yêu cầu mới:
Chuyển thông tin này thành một báo cáo in đen trắng với:
- Phân trang
- Tiêu đề
- Số âm đặt trong dấu ngoặc đơn

#### Thách thức kiến trúc:
Rõ ràng, mã mới cần được viết để xử lý việc in ấn. Tuy nhiên, thách thức kiến trúc nằm ở việc xác định: **liệu bao nhiêu mã cũ liên quan đến logic tính toán tài chính và logic trình bày trên web sẽ phải thay đổi?**

> Một kiến trúc phần mềm tốt phải giảm thiểu lượng mã bị thay đổi xuống mức tối thiểu, **lý tưởng là bằng không**.

### 3.3. Áp dụng OCP ở Cấp độ Kiến trúc

Để tuân thủ OCP, cần phân tách các vùng thay đổi (được xác định bởi SRP) và sau đó tổ chức sự phụ thuộc giữa chúng một cách hợp lý (sử dụng DIP).

#### Tách biệt Trách nhiệm

Về mặt khái niệm, việc tạo báo cáo bao gồm hai trách nhiệm chính:

1. **Tính toán dữ liệu** được báo cáo (_Financial Analyzer / Interactor_)
2. **Trình bày dữ liệu** đó dưới dạng thân thiện với môi trường cụ thể (_Web Reporter / Print Reporter_)

#### Thiết lập Phụ thuộc Unidirectional

> **Nguyên tắc tổ chức OCP ở cấp độ kiến trúc:** Nếu thành phần A (cao cấp, ổn định) cần được bảo vệ khỏi các thay đổi trong thành phần B (thấp cấp, dễ thay đổi), thì thành phần B phải phụ thuộc vào thành phần A.

Trong mô hình kiến trúc, **Interactor** (chứa logic nghiệp vụ tính toán) được đặt ở vị trí tuân thủ OCP tốt nhất. Các thành phần cấp thấp hơn, như:

- Database
- Controller  
- Presenters (_Screen Presenter, Print Presenter_)
- Views (_Web View, PDF View_)

Đều phải phụ thuộc vào **Interactor** hoặc các **Abstraction** của nó.

Nhờ cấu trúc phân cấp phụ thuộc đơn hướng này, các thay đổi trong giao diện người dùng (Views) hay cơ chế lưu trữ (Database) sẽ không gây ảnh hưởng đến logic nghiệp vụ cốt lõi bên trong Interactor.

> **OCP** là động lực thúc đẩy sự phân vùng hệ thống, bảo vệ các thành phần cấp cao khỏi các thay đổi trong các thành phần cấp thấp.

### 3.4. Kỹ thuật Triển khai OCP: Mẫu Thiết kế Strategy

Ở cấp độ lớp, OCP có thể được thực hiện thông qua **kế thừa**, hoặc sử dụng các **mẫu thiết kế hành vi** như **Strategy Pattern**.

**Strategy Pattern** là một mẫu thiết kế hành vi cho phép định nghĩa một họ thuật toán, đặt mỗi thuật toán vào một lớp riêng biệt (_ConcreteStrategy_), và làm cho chúng có thể hoán đổi cho nhau.

#### Ứng dụng của Strategy Pattern

```java
// Interface Strategy
public interface SortStrategy {
    void sort(List<Integer> list);
}

// Concrete Strategies
public class QuickSort implements SortStrategy {
    public void sort(List<Integer> list) { /* QuickSort implementation */ }
}

public class MergeSort implements SortStrategy {
    public void sort(List<Integer> list) { /* MergeSort implementation */ }
}

public class SelectionSort implements SortStrategy {
    public void sort(List<Integer> list) { /* SelectionSort implementation */ }
}

// Context class
public class SortedList {
    private SortStrategy strategy;
    
    public void setStrategy(SortStrategy strategy) {
        this.strategy = strategy;
    }
    
    public void sort(List<Integer> list) {
        strategy.sort(list);
    }
}
```

**Lớp Context** (ví dụ: `SortedList`) không bị sửa đổi khi một thuật toán mới được thêm vào. Thay vào đó, nó phụ thuộc vào một giao diện (`SortStrategy`). 

Các thuật toán cụ thể như `QuickSort`, `MergeSort`, hoặc `SelectionSort` là các triển khai của giao diện này. Khi cần thêm một thuật toán sắp xếp mới (ví dụ: `BubbleSort`), nhà phát triển chỉ cần tạo một `ConcreteStrategy` mới mà không cần chạm vào mã nguồn của `SortedList`.

> Điều này đáp ứng hoàn hảo yêu cầu **"mở để mở rộng, đóng để sửa đổi"**.   

---

## 4. Nguyên tắc Thay thế Liskov (Liskov Substitution Principle - LSP)

### 4.1. Định nghĩa Cốt lõi: Tính Thay thế Hành vi

**Nguyên tắc Thay thế Liskov (LSP)** là một nền tảng của lập trình hướng đối tượng, tuyên bố rằng:

> Các đối tượng của lớp con (_Subtype_) phải có khả năng thay thế đối tượng của lớp cha (_Base Type_) mà không làm thay đổi tính đúng đắn của chương trình.

**Nguyên tắc này** đảm bảo rằng các lớp con duy trì hành vi của lớp cha và không thay đổi logic khi được sử dụng để thay thế. 

Mặc dù ban đầu được hình thành để hướng dẫn sử dụng kế thừa, **LSP** đã mở rộng thành một nguyên tắc thiết kế rộng hơn áp dụng cho tất cả các mối quan hệ giao diện và triển khai.

### 4.2. Phân tích Trường hợp Vi phạm: Square và Rectangle

Trường hợp kinh điển vi phạm LSP là việc tạo lớp `Square` kế thừa từ lớp `Rectangle`.

#### Lớp Rectangle:
```java
public class Rectangle {
    protected int width;
    protected int height;
    
    public void setWidth(int width) {
        this.width = width;
    }
    
    public void setHeight(int height) {
        this.height = height;
    }
    
    public int getArea() {
        return width * height;
    }
}
```

Cho phép chiều rộng (_width_) và chiều cao (_height_) được thay đổi độc lập thông qua các phương thức `setWidth(int)` và `setHeight(int)`.

#### Lớp Square:
```java
public class Square extends Rectangle {
    @Override
    public void setWidth(int width) {
        this.width = width;
        this.height = width; // Bắt buộc phải bằng nhau
    }
    
    @Override
    public void setHeight(int height) {
        this.height = height;
        this.width = height; // Bắt buộc phải bằng nhau
    }
}
```

Do bản chất hình học, `Square` yêu cầu chiều rộng và chiều cao luôn bằng nhau. Khi triển khai, nếu `Square` ghi đè `setWidth(w)` hoặc `setHeight(h)`, nó buộc phải đặt cả hai thuộc tính này bằng giá trị mới truyền vào.

#### Sự phá vỡ Hợp đồng Hành vi

Khi một **Client** (_User_) tương tác với một đối tượng được khai báo là `Rectangle` nhưng thực chất là `Square`, Client mong đợi rằng nó có thể thay đổi chiều rộng và chiều cao một cách độc lập (tính chất của Rectangle).

```java
// Client code
Rectangle rect = new Square(); // Vi phạm LSP

rect.setWidth(5);
rect.setHeight(10);
int area = rect.getArea(); // Client mong đợi: 5×10=50
```

**Kết quả thực tế:**
- `rect.setWidth(5)` → width=5, height=5 (do Square)
- `rect.setHeight(10)` → width=10, height=10 (do Square)  
- `rect.getArea()` → 10×10=100

**Mong đợi của Client:** 5×10=50  
**Thực tế:** 10×10=100

> Sự sai lệch này phá vỡ logic của Client và chứng tỏ `Square` không phải là một _Subtype_ hợp lệ của `Rectangle` theo LSP.

### 4.3. Tác động Kiến trúc của LSP

**LSP** là một nguyên tắc quan trọng để duy trì tính toàn vẹn của hệ thống. Khi LSP bị vi phạm, lập trình viên buộc phải thêm các cơ chế bổ sung để kiểm tra và quản lý hành vi đặc biệt của các lớp con. Điều này thường dẫn đến việc sử dụng mã kiểm tra kiểu (_type checking_) hoặc các toán tử như `instanceof`.

#### Ví dụ về vi phạm LSP trong kiến trúc:

```java
public void processShape(Rectangle rect) {
    // Vi phạm LSP - phải kiểm tra kiểu
    if (rect instanceof Square) {
        // Logic đặc biệt cho Square
        rect.setWidth(10);
    } else {
        // Logic cho Rectangle
        rect.setWidth(10);
        rect.setHeight(20);
    }
}
```

**Vấn đề:** Việc phải viết mã kiểm tra kiểu để xác định đối tượng đang được sử dụng là gì để thay đổi hành vi tương ứng sẽ làm **ô nhiễm kiến trúc**. 

Nó làm giảm tính linh hoạt, vì việc thêm một lớp con mới đòi hỏi phải sửa đổi tất cả các đoạn mã kiểm tra kiểu hiện có.

> **LSP** đảm bảo rằng tính đa hình được sử dụng một cách đáng tin cậy, giúp hệ thống hoạt động đồng nhất và dễ bảo trì hơn.   

---

## 5. Nguyên tắc Phân tách Giao diện (Interface Segregation Principle - ISP)

### 5.1. Định nghĩa Cốt lõi và Vấn đề "Fat Interface"

**Nguyên tắc Phân tách Giao diện (ISP)** khẳng định rằng:

> **Khách hàng (_Client_) không nên bị buộc phải phụ thuộc vào các giao diện (_interface_) hoặc các phương thức mà nó không sử dụng.**

**Mục tiêu của ISP** là tránh các **"Fat Interface"** (_giao diện khổng lồ_) chứa nhiều phương thức không liên quan. Thay vào đó, giao diện nên nhỏ, chuyên biệt và tập trung vào vai trò cụ thể.

### 5.2. Phân tích Trường hợp Vi phạm

#### Ví dụ: Interface Animal vi phạm ISP

```java
// Fat Interface - Vi phạm ISP
public interface Animal {
    void eat();
    void drink();
    void sleep();
    void swim();    // Không phải tất cả động vật đều biết bơi
    void fly();     // Không phải tất cả động vật đều biết bay
}
```

Xét một interface chung `Animal` chứa các phương thức: `void Eat()`, `void Drink()`, `void Sleep()`, `void Swim()`, và `void Fly()`.

#### Vấn đề khi triển khai:

```java
public class Cat implements Animal {
    public void eat() { /* Cat eats */ }
    public void drink() { /* Cat drinks */ }
    public void sleep() { /* Cat sleeps */ }
    
    // Cat không biết bơi - phải triển khai như thế nào?
    public void swim() {
        throw new UnsupportedOperationException("Cats can't swim!");
    }
    
    // Cat không biết bay - phải triển khai như thế nào?
    public void fly() {
        throw new UnsupportedOperationException("Cats can't fly!");
    }
}
```

Nếu một lớp như `Cat` triển khai interface `Animal`, nó sẽ bị buộc phải cung cấp các triển khai cho cả `Swim()` và `Fly()`. Do `Cat` về mặt sinh học không thể bay, nhà phát triển sẽ phải triển khai các phương thức này bằng cách để trống hoặc ném ra ngoại lệ. 

> Việc này tạo ra sự phức tạp không cần thiết, mơ hồ về chức năng của lớp, và vi phạm ISP.

### 5.3. Giải pháp Phân tách theo Vai trò

Giải pháp để tuân thủ ISP là phân tách interface lớn thành các giao diện nhỏ hơn, chuyên biệt theo vai trò của Client.

#### Cách tiếp cận đúng:

```java
// Interface cơ bản
public interface Animal {
    void eat();
    void drink();
    void sleep();
}

// Interface chuyên biệt cho động vật biết bay
public interface Flyable {
    void fly();
}

// Interface chuyên biệt cho động vật biết bơi
public interface Swimmable {
    void swim();
}

// Interface chuyên biệt cho động vật có thể chạy
public interface Runnable {
    void run();
}
```

#### Triển khai tuân thủ ISP:

```java
public class Cat implements Animal, Runnable {
    public void eat() { /* Cat eats */ }
    public void drink() { /* Cat drinks */ }
    public void sleep() { /* Cat sleeps */ }
    public void run() { /* Cat runs */ }
    // Không cần triển khai fly() hay swim()
}

public class Bird implements Animal, Flyable {
    public void eat() { /* Bird eats */ }
    public void drink() { /* Bird drinks */ }
    public void sleep() { /* Bird sleeps */ }
    public void fly() { /* Bird flies */ }
    // Không cần triển khai swim()
}

public class Duck implements Animal, Flyable, Swimmable, Runnable {
    public void eat() { /* Duck eats */ }
    public void drink() { /* Duck drinks */ }
    public void sleep() { /* Duck sleeps */ }
    public void fly() { /* Duck flies */ }
    public void swim() { /* Duck swims */ }
    public void run() { /* Duck runs */ }
}
```

**Kết quả:**
- Lớp `Cat` chỉ cần triển khai `Animal` và `Runnable`
- Lớp `Bird` triển khai `Animal` và `Flyable`
- Lớp `Duck` triển khai tất cả các interface cần thiết

> Bằng cách này, mỗi lớp không bị buộc phải phụ thuộc vào các phương thức mà nó không sử dụng, giảm thiểu sự phức tạp và tăng tính rõ ràng của mã nguồn.

### 5.4. Tác động Kiến trúc và Quản lý Biến động

**ISP** có ý nghĩa sâu sắc ở cấp độ kiến trúc, đặc biệt trong các ngôn ngữ kiểu tĩnh như Java. Các khai báo interface trong các ngôn ngữ này tạo ra sự phụ thuộc mã nguồn cứng nhắc.

#### Vấn đề Phụ thuộc vào "Hành lý"

```
System (S) → Framework (F) → Database (D)
```

Khi một hệ thống (S) phụ thuộc vào một framework (F), và framework này lại phụ thuộc vào một cơ sở dữ liệu (D), nếu D chứa các tính năng mà F và S không sử dụng (gọi là **"hành lý"**), thì một thay đổi đối với những tính năng không sử dụng đó trong D vẫn có thể buộc F và S phải tái biên dịch hoặc tái triển khai.

> Tình trạng này làm tăng chi phí bảo trì và giảm sự ổn định của hệ thống.

**ISP** giải quyết vấn đề này bằng cách đảm bảo rằng các module chỉ phụ thuộc vào các interface tối thiểu cần thiết. Bằng cách giảm kích thước và sự chuyên biệt của interface, ISP trực tiếp giảm thiểu nguy cơ xảy ra sự biến động lan truyền từ các module xa xôi, không liên quan.

---

## 6. Nguyên tắc Đảo ngược Sự phụ thuộc (Dependency Inversion Principle - DIP)

### 6.1. Định nghĩa Cốt lõi: Phụ thuộc vào Abstraction

**Nguyên tắc Đảo ngược Sự phụ thuộc (DIP)** là nguyên tắc tổ chức quan trọng nhất ở cấp độ kiến trúc, tuyên bố rằng:

> Các hệ thống linh hoạt nhất là những hệ thống trong đó sự phụ thuộc mã nguồn chỉ hướng tới các **abstraction** (_trừu tượng_), không phải các **concretion** (_cụ thể_).

**DIP** được xác định thông qua **hai tuyên bố chính**:

1. **Module cấp cao** (_High-level modules_) không nên phụ thuộc vào **module cấp thấp** (_Low-level modules_). Cả hai nên phụ thuộc vào **abstractions**.

2. **Abstractions** không nên phụ thuộc vào **details** (_chi tiết_). **Details** nên phụ thuộc vào **abstractions**.

**Module cấp cao** chứa chính sách nghiệp vụ quan trọng (_logic cốt lõi_), trong khi **module cấp thấp** chứa các chi tiết triển khai (_như tương tác cơ sở dữ liệu, giao tiếp API_). 

> **DIP** đảm bảo module cấp cao được cách ly khỏi sự biến động của module cấp thấp.

### 6.2. Cơ chế Đảo ngược Sự phụ thuộc (Dependency Inversion)

#### Kiến trúc truyền thống (vi phạm DIP):

```
High-Level Module (HL) → Low-Level Module (LL)
         ↓                        ↓
    Logic nghiệp vụ         Chi tiết triển khai
         ↓                        ↓
    Phụ thuộc vào LL         Database, API, etc.
```

Trong kiến trúc phần mềm truyền thống, **luồng kiểm soát** (_Flow of Control_) đi từ module cấp cao (HL) xuống module cấp thấp (LL), và **sự phụ thuộc mã nguồn** đi theo cùng một hướng.

#### Kiến trúc với DIP:

```
High-Level Module (HL) ←→ Abstraction (Interface)
         ↓                        ↑
    Logic nghiệp vụ         Low-Level Module (LL)
                                ↓
                         Chi tiết triển khai
```

**DIP** đảo ngược hướng phụ thuộc mã nguồn. Bằng cách chèn một **interface** (_Abstraction_) giữa HL và LL:

- **Module cấp cao (HL)** định nghĩa interface
- **Module cấp thấp (LL)** triển khai interface đó

#### Hai luồng khác nhau:

1. **Luồng Kiểm soát** (_Runtime_): Vẫn đi từ HL sang LL
2. **Sự phụ thuộc Mã nguồn** (_Compile-time_): Hướng từ LL lên Interface, sau đó Interface được HL sử dụng

> Do LL phụ thuộc vào Interface mà HL định nghĩa, hướng phụ thuộc đã bị đảo ngược so với luồng kiểm soát.

**Sự đảo ngược này** là chìa khóa để đạt được khả năng mở rộng (OCP) và khả năng kiểm thử cao. Module cấp cao có thể được kiểm thử độc lập bằng cách thay thế các triển khai cụ thể (LL) bằng các đối tượng giả (_mock/stub_).

### 6.3. Khái niệm Abstraction Ổn định và Thực tiễn Mã hóa Tốt

Các **interface** (_abstractions_) thường ít biến động hơn so với các **lớp triển khai cụ thể** (_concretions_). Thay đổi đối với một lớp triển khai hiếm khi yêu cầu thay đổi interface mà nó triển khai. 

> Do đó, các nhà thiết kế phần mềm luôn cố gắng giảm thiểu sự biến động của các interface.

#### Để duy trì tính ổn định, cần tuân thủ các thực tiễn mã hóa nghiêm ngặt:

- ❌ **Không tham chiếu** đến các lớp cụ thể và dễ thay đổi
- ❌ **Không kế thừa** từ các lớp cụ thể và dễ thay đổi  
- ❌ **Không ghi đè** các hàm cụ thể (_concrete functions_)
- ❌ **Tuyệt đối không đề cập** đến tên của bất kỳ thứ gì cụ thể và dễ thay đổi trong các module cấp cao

### 6.4. Kỹ thuật Quản lý Concretion: Sử dụng Mẫu Factory

Trong hầu hết các ngôn ngữ hướng đối tượng, việc tạo đối tượng cụ thể (ví dụ: `new ConcreteObject()`) tạo ra sự phụ thuộc mã nguồn cứng nhắc vào lớp cụ thể đó. 

> Sự phụ thuộc này vi phạm các quy tắc của DIP.

Để giải quyết vấn đề này, các **Mẫu Factory** (_Abstract Factory_ hoặc _Factory Method_) được sử dụng để quản lý và cô lập sự phụ thuộc không mong muốn này.

#### Ví dụ về Factory Method Pattern:

```java
// Interface abstraction
public interface DatabaseConnection {
    void connect();
    void disconnect();
    void executeQuery(String query);
}

// Concrete implementations
public class OracleConnection implements DatabaseConnection {
    public void connect() { /* Oracle specific */ }
    public void disconnect() { /* Oracle specific */ }
    public void executeQuery(String query) { /* Oracle specific */ }
}

public class MySQLConnection implements DatabaseConnection {
    public void connect() { /* MySQL specific */ }
    public void disconnect() { /* MySQL specific */ }
    public void executeQuery(String query) { /* MySQL specific */ }
}

// Factory để quản lý việc tạo đối tượng
public class DatabaseConnectionFactory {
    public static DatabaseConnection createConnection(String type) {
        switch (type.toLowerCase()) {
            case "oracle":
                return new OracleConnection(); // Phụ thuộc cụ thể được cô lập ở đây
            case "mysql":
                return new MySQLConnection(); // Phụ thuộc cụ thể được cô lập ở đây
            default:
                throw new IllegalArgumentException("Unsupported database type");
        }
    }
}

// High-level module - chỉ phụ thuộc vào abstraction
public class UserService {
    private DatabaseConnection dbConnection;
    
    public UserService(String dbType) {
        // Phụ thuộc vào abstraction, không phụ thuộc vào implementation cụ thể
        this.dbConnection = DatabaseConnectionFactory.createConnection(dbType);
    }
    
    public void saveUser(User user) {
        dbConnection.connect();
        dbConnection.executeQuery("INSERT INTO users...");
        dbConnection.disconnect();
    }
}
```

**Factory Method** là một mẫu thiết kế khởi tạo cung cấp một interface để tạo đối tượng trong lớp cha (_Creator_), nhưng cho phép các lớp con (_ConcreteCreator_) xác định loại đối tượng cụ thể sẽ được tạo ra.

> Client chỉ tương tác với interface và Factory, cách ly nó khỏi chi tiết cụ thể về loại cơ sở dữ liệu được sử dụng. Sự phụ thuộc vào lớp cụ thể (như `new OracleConnection()`) bị giới hạn bên trong Factory.

### 6.5. Sự Mở rộng của DIP: Quy tắc Phụ thuộc (The Dependency Rule)

Ở cấp độ kiến trúc, **DIP** là nguyên tắc tổ chức quan trọng nhất. Nó định nghĩa **ranh giới kiến trúc** (_Architectural Boundary_) ngăn cách phần trừu tượng (_Abstraction_) khỏi phần cụ thể (_Concretion_).

#### Quy tắc Phụ thuộc (The Dependency Rule):

> **Tất cả các sự phụ thuộc mã nguồn phải băng qua ranh giới kiến trúc theo một hướng duy nhất, luôn hướng về phía các thực thể trừu tượng (phần lõi kiến trúc).**

```
┌─────────────────────────────────────┐
│           DETAILS                   │
│  (Database, UI, Framework, etc.)    │
└─────────────────┬───────────────────┘
                  │ Phụ thuộc hướng lên
                  ▼
┌─────────────────────────────────────┐
│         ABSTRACTIONS                │
│    (Business Logic, Core)           │
└─────────────────────────────────────┘
```

> Điều này đảm bảo rằng các module chứa logic nghiệp vụ cốt lõi (_abstraction_) không bao giờ bị phụ thuộc vào các chi tiết triển khai bên ngoài (_concretion_), từ đó duy trì tính ổn định, linh hoạt, và khả năng kiểm thử cao cho toàn bộ hệ thống.

---

## Tóm tắt toàn bộ SOLID

### Bảng tóm tắt 5 nguyên tắc SOLID:

| **Nguyên tắc** | **Tên viết tắt** | **Mục tiêu chính** | **Vấn đề giải quyết** |
|----------------|------------------|-------------------|---------------------|
| **Single Responsibility** | **S** | Một lớp chỉ có một lý do để thay đổi | Tách biệt trách nhiệm theo Actor |
| **Open-Closed** | **O** | Mở để mở rộng, đóng để sửa đổi | Thêm tính năng mới mà không sửa code cũ |
| **Liskov Substitution** | **L** | Lớp con có thể thay thế lớp cha | Đảm bảo tính đa hình đáng tin cậy |
| **Interface Segregation** | **I** | Client không phụ thuộc vào interface không sử dụng | Tránh Fat Interface |
| **Dependency Inversion** | **D** | Phụ thuộc vào abstraction, không phải concretion | Đảo ngược hướng phụ thuộc |

### Mối quan hệ giữa các nguyên tắc:

1. **SRP** là nền tảng - xác định ranh giới và trách nhiệm
2. **OCP** xây dựng trên SRP - tổ chức để dễ mở rộng
3. **LSP** đảm bảo tính nhất quán trong kế thừa
4. **ISP** tinh chỉnh interface để phù hợp với client
5. **DIP** là nguyên tắc tổ chức cao nhất - định hướng kiến trúc

### Lợi ích khi áp dụng SOLID:

- ✅ **Maintainability**: Dễ bảo trì và sửa lỗi
- ✅ **Testability**: Dễ dàng viết unit test
- ✅ **Flexibility**: Linh hoạt với thay đổi yêu cầu
- ✅ **Reusability**: Tái sử dụng component hiệu quả
- ✅ **Scalability**: Dễ dàng mở rộng hệ thống

### Kết luận:

**SOLID** không chỉ là 5 nguyên tắc riêng lẻ, mà là một **hệ thống nguyên tắc tích hợp** giúp tạo ra kiến trúc phần mềm bền vững, linh hoạt và có thể bảo trì. Việc áp dụng đúng SOLID sẽ giúp team phát triển phần mềm hiệu quả hơn và tạo ra sản phẩm chất lượng cao.

