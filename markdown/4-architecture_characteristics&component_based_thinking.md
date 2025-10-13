# Tư Duy Dựa Trên Thành Phần, Độ Hạt Kiến Trúc, và Các Chỉ Số Đo Lường Tính Mô-đun

## I. Nền Tảng Lý Thuyết Về Cấu Trúc Thành Phần

### I.1. Phân Biệt Đơn Vị Thực Thi (Module) và Đơn Vị Triển Khai (Component)

Trong thiết kế kiến trúc phần mềm, việc thiết lập các ranh giới lô-gíc và vật lý là điều kiện tiên quyết. **Module** và **Component**, mặc dù thường được sử dụng thay thế cho nhau, nhưng lại đại diện cho các khái niệm khác biệt ở cấp độ thiết kế và triển khai.

**Module** được định nghĩa là một tập hợp mã liên quan (_a collection of related code_) và đóng vai trò là **đơn vị thực thi** (_unit of implementation_). Module thường thể hiện các ranh giới lô-gíc trong mã nguồn, chẳng hạn như:

- **Gói** (_package_) trong Java
- **Namespace** trong .NET

Được nhà phát triển sử dụng để nhóm các chức năng hoặc lớp có liên quan.

**Component** là một thực thể thời gian chạy (_a runtime entity_), định nghĩa một **đơn vị triển khai** (_unit of deployment_). Component hoạt động trong không gian địa chỉ riêng hoặc chung, có khả năng được triển khai độc lập.

#### Các Loại Component:

1. **Thư viện** (_Library_)
   - Chạy trong cùng một không gian bộ nhớ với mã gọi
   - Giao tiếp thông qua các lời gọi hàm ngôn ngữ

2. **Dịch vụ** (_Service_)
   - Chạy trong không gian địa chỉ riêng của nó
   - Giao tiếp qua các giao thức mạng cấp thấp hoặc các định dạng cấp cao như **REST** hoặc **message queue**

> **Sự phân biệt này cực kỳ quan trọng:** Một hệ thống **Client-Server** có thể chỉ có hai module lô-gíc (client và server) nhưng lại phân rã thành hàng chục component vật lý (ví dụ: các dịch vụ phân tán, cơ sở dữ liệu, bộ xử lý sự kiện).

### I.2. Tầm Quan Trọng Của Độ Hạt Thành Phần (Component Granularity)

**Độ hạt** (_Granularity_) là quyết định về kích thước và phạm vi trách nhiệm của một component, và nó quyết định cách các **Đặc tính Kiến trúc** (_Architecture Characteristics - ACs_) được hiện thực hóa.

> Việc chọn độ hạt không chính xác là một quyết định đánh đổi cấp cao (_trade-off_) trong kiến trúc.

#### **Component Quá Mịn (Too Fine-Grained)**

**Đặc điểm:**
- Component quá nhỏ hoặc quá chuyên biệt
- Dẫn đến quá nhiều giao tiếp giữa các component (_too much communication_) để hoàn thành các tác vụ nghiệp vụ

**Tác động tiêu cực:**
- Tăng **độ trễ** (_latency_)
- Tăng **chi phí vận hành** (_overhead_)
- Ảnh hưởng trực tiếp và tiêu cực đến đặc tính **Performance** của hệ thống

#### **Component Quá Thô (Too Coarse-Grained)**

**Đặc điểm:**
- Component quá lớn
- Khuyến khích khớp nối nội bộ cao (_high internal coupling_)

**Tác động tiêu cực:**
- Gây khó khăn đáng kể cho việc **triển khai độc lập** (_Deployability_)
- Giảm khả năng **kiểm thử** (_Testability_)
- Làm giảm tính **mô-đun** (_Modularity_) của hệ thống
- Thường là dấu hiệu của việc vi phạm **Nguyên tắc Trách nhiệm Đơn lẻ** (**SRP**)

> Khi một component đảm nhận quá nhiều trách nhiệm, nó có **tính gắn kết thấp** (_Low Cohesion_), dẫn đến việc một thay đổi chức năng có thể gây ra **hiệu ứng lan truyền** (_ripple effect_) trong chính component đó, làm giảm đáng kể khả năng **bảo trì** (_Maintainability_).

---

## II. Chu Trình và Phương Pháp Nhận Dạng Thành Phần Lô-gíc

Việc nhận dạng các component ban đầu hiếm khi tạo ra thiết kế tối ưu ngay lập tức. Do đó, kiến trúc sư cần phải áp dụng một **phương pháp tiếp cận lặp** (_iterative approach_) thông qua **Chu trình Nhận dạng Thành phần** (_Component Identification Flow_) để liên tục tinh chỉnh ranh giới component.

### Chu Trình 5 Bước:

1. **Xác định Thành phần Ban đầu** (_Identify Initial Components_)
   - Quyết định các component cấp cao để bắt đầu
   - Thường dựa trên phân vùng theo miền (_domain_) hoặc theo kỹ thuật

2. **Gán Yêu cầu cho Thành phần** (_Assign Requirements to Components_)
   - Ánh xạ các yêu cầu chức năng (_functional requirements_) vào các component ban đầu
   - Nếu một component ôm đồm quá nhiều trách nhiệm, bước này yêu cầu tạo mới, hợp nhất, hoặc chia nhỏ component

3. **Phân tích Vai trò và Trách nhiệm** (_Analyze Roles and Responsibilities_)
   - Kiểm tra các vai trò và hành vi được xác định trong quá trình phân tích yêu cầu
   - Đảm bảo độ hạt của component phù hợp với ngữ cảnh miền

4. **Phân tích Đặc tính Kiến trúc** (_Analyze Architecture Characteristics_)
   - Đánh giá cách các **ACs** đã được ưu tiên (ví dụ: **Scalability**, **Security**) ảnh hưởng đến sự phân chia và độ hạt của component
   - Ví dụ: Yêu cầu về tính sẵn sàng (_Availability_) cao có thể đòi hỏi component đó phải được cô lập và triển khai độc lập

5. **Tái cấu trúc Thành phần** (_Restructure Components_)
   - Quá trình lặp lại và tinh chỉnh thiết kế component
   - Thường được thực hiện với sự cộng tác của nhà phát triển

### II.1. Chống Mẫu (Anti-Pattern): Entity Trap

Một lỗi thiết kế nghiêm trọng, thường xảy ra ở bước khởi tạo component, là **Entity Trap** (_Bẫy Thực thể_).

**Định nghĩa:** Chống mẫu này xuất hiện khi kiến trúc sư lấy mỗi thực thể (_entity_) xác định trong yêu cầu (ví dụ: Customer, Ticket) và tạo ra một component Manager tương ứng (ví dụ: Customer Manager, Ticket Manager).

**Vấn đề:**
- Thiết kế dựa trên Entity Trap không phải là một kiến trúc tập trung vào hành vi mà là một ánh xạ quan hệ của framework vào cơ sở dữ liệu
- Các component này thường vi phạm **SRP** vì chúng trở thành **"bãi rác"** (_dumping ground_) cho các chức năng không liên quan
- Khiến chúng trở nên quá lớn và mất đi mục đích rõ ràng

**Ví dụ:** Ticket Manager có thể xử lý cả:
- `create ticket`
- `assign ticket to expert`
- `send out survey`

### II.2. Các Phương Pháp Nhận Dạng Thành Phần Lô-gíc Hiệu Quả

Để tạo ra các component có tính gắn kết cao, kiến trúc sư nên tập trung vào **hành vi** và **trách nhiệm**.

#### **Phương Pháp Actor/Actions**

**Định nghĩa:** Phương pháp này, được xác định bởi **Rational Unified Process**, tập trung vào việc xác định các **tác nhân** (_actors_) và các **hành động** (_actions_) mà họ thực hiện với ứng dụng.

**Quy trình:**
- Bằng cách nhóm các hành động liên quan, kiến trúc sư có thể xác định **ranh giới miền** (_Domain Boundaries_) rõ ràng
- Ví dụ: Nhóm các hành động liên quan đến Customer thành **Customer Registration** và các hành động liên quan đến Ticket thành **Ticket Creation**

**Kết quả:**
- Tạo ra các component có **tính gắn kết chức năng** (_Functional Cohesion_) cao
- Phù hợp để phân vùng theo trách nhiệm nghiệp vụ

#### **Phương pháp Workflow**

**Định nghĩa:** Phương pháp **Workflow** (_Luồng Công việc_) mô hình hóa các component dựa trên các quy trình nghiệp vụ tuần tự.

**Áp dụng:**
- Hữu ích khi kiến trúc sư có kiến thức cơ bản về các luồng người dùng chính trong hệ thống
- Các component được thiết kế để xử lý một chuỗi bước cụ thể trong luồng công việc

**Ví dụ:** `Register With Site → Create Ticket → Assign Ticket`

**Lưu ý:**
- Mặc dù tạo ra sự **gắn kết tuần tự** cao, cần cẩn thận để tránh khớp nối quá mức nếu các luồng công việc phức tạp hoặc giao nhau

---

## III. Đo Lường và Tối Ưu Hóa Tính Mô-đun (Modularity)

Tính mô-đun được đo lường bằng cách định lượng hóa hai đặc điểm:
- **Tính Gắn kết** (_Cohesion_)
- **Khớp nối** (_Coupling_)

### III.1. Đánh Giá Tính Gắn Kết (Cohesion)

**Tính gắn kết** (_Cohesion_) đo lường mức độ các phần của một module có liên quan đến nhau và đóng gói chức năng đó.

#### **Đặc điểm:**

**Gắn kết cao** là lý tưởng vì nó:
- Thúc đẩy tính **đóng gói** (_encapsulation_)
- Giảm khả năng xảy ra lỗi

**Gắn kết thấp** cho thấy:
- Thiết kế kém
- Cần phải tái cấu trúc

#### **Các Loại Gắn Kết:**

1. **Gắn kết Chức năng** (_Functional Cohesion_) - **Cao nhất**
   - Module chứa mọi thứ cần thiết để thực hiện một chức năng duy nhất

2. **Gắn kết Ngẫu nhiên** (_Coincidental Cohesion_) - **Tệ nhất**
   - Các phần tử không có mối liên hệ nào ngoài việc nằm chung trong một file nguồn

#### **Đo Lường Định Lượng: LCOM**

Để định lượng tính gắn kết ở cấp độ lớp, người ta sử dụng **Chỉ số Thiếu Gắn kết trong Phương thức** (_LCOM - Lack of Cohesion in Methods_).

**Chỉ số LCOM2** được tính bằng:
```
LCOM2 = P - Q
```

**Trong đó:**
- **P** = Số cặp phương thức **không chia sẻ** thuộc tính
- **Q** = Số cặp phương thức **chia sẻ** thuộc tính

> **Nếu LCOM2 có giá trị cao**, điều đó báo hiệu tính gắn kết thấp và gợi ý rằng lớp nên được tách thành các đơn vị nhỏ hơn.

### III.2. Đánh Giá Khớp Nối (Coupling) và Tính Bất Ổn (Instability)

**Khớp nối** (_Coupling_) là mức độ phụ thuộc lẫn nhau giữa các module. Khớp nối được định lượng thông qua hai chỉ số chính:

#### **1. Afferent Coupling (Cₐ)**

- **Khớp nối hướng vào** (_incoming dependency_)
- **Cₐ cao** cho thấy component là một trung tâm ổn định, có nhiều thành phần bên ngoài phụ thuộc vào nó
- **Giá trị khuyến nghị:** 0 → 500

#### **2. Efferent Coupling (Cₑ)**

- **Khớp nối hướng ra** (_outgoing dependency_)
- **Cₑ cao** (lớn hơn 20) cho thấy component rất dễ bị ảnh hưởng bởi thay đổi của các component bên ngoài mà nó phụ thuộc

#### **Chỉ Số Instability (I)**

**Chỉ số Instability** được sử dụng để xác định tính biến động (_volatility_) của một component, được tính bằng tỷ lệ giữa Cₑ và tổng coupling:

```
I = Cₑ / (Cₑ + Cₐ)
```

> Chỉ số Instability cung cấp một cái nhìn định lượng về việc component có tuân thủ **Nguyên tắc Đảo ngược Sự phụ thuộc** (**DIP**) hay không.

#### **Bảng Phân Tích Instability Index**

| **Phạm vi I (Instability)** | **Đặc điểm Về Phụ thuộc** | **Tính Biến động & Vai trò Kiến trúc** | **Nguyên tắc Liên quan** |
|----------------------------|---------------------------|--------------------------------------|-------------------------|
| **I≈0 (Rất Ổn định) (0.0 – 0.3)** | **Cₐ cao**, **Cₑ thấp** | **Module Core**, **Abstraction**, **Interfaces**. Được bảo vệ khỏi sự thay đổi (_Protected from change_). | Tuân thủ **DIP** (High-Level Module), **OCP** (Closed for Modification). |
| **I∈(0.3,0.7) (Ổn định Trung bình)** | **Cₐ ≈ Cₑ** | **NÊN TRÁNH**. Vừa dễ thay đổi, vừa gây sự cố lan truyền (_Ripple Effects_). | Vi phạm nguyên tắc phân tầng rõ ràng của kiến trúc. |
| **I≈1 (Rất Bất ổn) (0.7 – 1.0)** | **Cₑ cao**, **Cₐ thấp** | **Module Concrete**, **Chi tiết cấp thấp** (_Low-Level Detail_). Được mong đợi sẽ thay đổi thường xuyên. | Tuân thủ **DIP** (Low-Level Detail). |

> Việc đảm bảo các module cốt lõi (ví dụ: các giao diện nghiệp vụ) nằm trong vùng ổn định (**I≈0**) là cách **DIP** đảm bảo rằng các module cấp cao này có thể tuân thủ **Nguyên tắc Mở/Đóng** (**OCP**), nghĩa là chúng có thể được mở rộng mà không cần sửa đổi khi các chi tiết triển khai cấp thấp thay đổi.

---

## IV. Chuyển Đổi Kiến Trúc Lô-gíc Sang Vật Lý và Ra Quyết Định

### IV.1. Architecture Quantum: Đơn Vị Triển Khai Độc Lập

**Architecture Quantum** được định nghĩa là một **đơn vị triển khai độc lập** (_an independently deployable artifact_) có:
- **Tính gắn kết chức năng cao**
- **Khớp nối động đồng bộ**

**Đặc điểm:**
- **Quantum** là đơn vị cơ bản trong các kiến trúc phân tán (ví dụ: **Microservices**)
- Thiết lập ranh giới vật lý
- Hiểu rõ Quantum là điều kiện cần thiết để thiết kế các hệ thống có:
  - **Tính sẵn sàng** (_Availability_) cao
  - **Khả năng phục hồi** (_Recoverability_) cao

### IV.2. Ánh Xạ Kiến Trúc Lô-gíc Sang Vật Lý

Quá trình ánh xạ các component lô-gíc (_Domains_) sang môi trường vật lý (ví dụ: cloud infrastructure) phải được định hướng bởi các **Đặc tính Kiến trúc ưu tiên**.

#### **Ví dụ: Scalability là Ưu Tiên Hàng Đầu**

**Kiến trúc vật lý sẽ ưu tiên:**
- Các cơ chế mở rộng ngang (_horizontal scaling_)
- Sử dụng **Elastic Load Balancer** (**ELB**) để phân phối tải trên nhiều Instance máy chủ ứng dụng

#### **Chiến Lược Cơ Sở Dữ Liệu Vật Lý (Polyglot Persistence)**

Việc phân chia component lô-gíc thành các Quantum cũng định hình chiến lược cơ sở dữ liệu vật lý.

**Các loại cơ sở dữ liệu khác nhau** được sử dụng để tối ưu hóa các **ACs** cụ thể:

| **Cơ Sở Dữ Liệu** | **Mục Đích Sử Dụng** | **AC Được Tối Ưu** |
|-------------------|---------------------|-------------------|
| **Redis** | Caching | **Performance** |
| **MySQL** | Dữ liệu chính | **Tính toàn vẹn dữ liệu** |
| **MongoDB** | Dữ liệu phi cấu trúc | **Flexibility** |
| **Cassandra** | Dữ liệu phân tán | **Scalability** |

### IV.3. Lựa Chọn Phong Cách Kiến Trúc (Architecture Style)

Việc chọn **Phong cách Kiến trúc** (**Layered**, **Modular Monolith**, **Microservices**, **Event-Driven**, v.v.) là một quyết định chiến lược để định hình **Kiến trúc Vật lý**.

> Quyết định này được thực hiện thông qua việc phân tích đánh đổi trên **Architecture Styles Worksheet**.

#### **Đánh Đổi Quan Trọng: Simplicity vs Scalability**

**Một trong những đánh đổi quan trọng nhất** là giữa tính đơn giản và khả năng mở rộng.

**Các phong cách kiến trúc được thiết kế để tối ưu hóa Scalability và Elasticity** (như **Microservices**):
- **Ưu điểm:** **Scalability** và **Elasticity** cao
- **Nhược điểm:** 
  - Chi phí phát triển và vận hành cao hơn (_Cost_)
  - Độ phức tạp cao hơn, làm giảm **Simplicity**

**Các kiến trúc đơn giản hơn** (như **Layered**):
- **Ưu điểm:** **Simplicity** cao
- **Nhược điểm:** Gặp thách thức lớn hơn khi cần đạt **Scalability** cực cao

### IV.4. Triết Lý Đánh Đổi và "Least Worst Architecture"

> **Nguyên tắc vàng trong kiến trúc:** "Mọi thứ đều là sự đánh đổi" (_"Everything is a trade-off"_).

Vì không có kiến trúc nào hoàn hảo, mục tiêu của kiến trúc sư là đạt được **kiến trúc ít tệ nhất** (_the least worst architecture_), tức là kiến trúc:
- **Tối ưu hóa** các **Đặc tính Kiến trúc** quan trọng nhất
- **Chấp nhận** những khuyết điểm có thể chấp nhận được ở các **AC** ít quan trọng hơn

#### **Ví Dụ Đánh Đổi: Security vs Performance**

**Tình huống:** Việc tăng cường bảo mật thông qua **mã hóa on-the-fly** (_on-the-fly encryption_)

**Tác động:**
- **Tăng cường:** **Security** ✅
- **Giảm thiểu:** **Performance** ❌ (tăng độ phức tạp và chi phí xử lý)

#### **Phân Tích Đánh Đổi Hệ Thống Đấu Giá (Queues vs. Topics)**

Việc lựa chọn giữa **Queues** (_Point-to-Point_) và **Topics** (_Publish-and-Subscribe_) trong hệ thống nhắn tin là một minh chứng cho tư duy đánh đổi.

**Nếu Agility** (khả năng thêm người tiêu thụ mới nhanh chóng) là quan trọng:
- **Topics** là lựa chọn ưu việt
- **Lý do:** Bid Producer không cần biết về Consumer (**Khớp nối thấp**)

**Nếu Observability** (giám sát từng hàng đợi riêng biệt) và **Extensibility** (hỗ trợ hợp đồng dữ liệu riêng biệt cho từng dịch vụ) là ưu tiên:
- **Queues** lại là giải pháp tốt hơn

---

## Kết Luận

Tư duy dựa trên thành phần đòi hỏi kiến trúc sư phải:

1. **Phân biệt rõ ràng** giữa Module (đơn vị thực thi) và Component (đơn vị triển khai)
2. **Áp dụng chu trình lặp** để nhận dạng và tinh chỉnh component
3. **Tránh các chống mẫu** như Entity Trap
4. **Sử dụng các chỉ số định lượng** (LCOM, Coupling, Instability) để đánh giá chất lượng
5. **Áp dụng triết lý đánh đổi** để đạt được kiến trúc "ít tệ nhất" phù hợp với ngữ cảnh cụ thể

Việc hiểu rõ và áp dụng các nguyên tắc này sẽ giúp kiến trúc sư thiết kế ra các hệ thống có tính mô-đun cao, dễ bảo trì và có thể thích ứng với sự thay đổi.