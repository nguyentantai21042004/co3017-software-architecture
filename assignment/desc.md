# Hệ thống Gia sư Thông minh (Intelligent Tutoring System - ITS)

## Mục tiêu

Mục đích của bài tập này là thiết kế một kiến trúc phần mềm và triển khai Hệ thống Gia sư Thông minh (ITS) sử dụng các nguyên tắc SOLID. Hệ thống này cung cấp trải nghiệm học tập được cá nhân hóa và thích ứng cho học sinh/sinh viên. ITS phải có khả năng đánh giá kiến thức của học sinh/sinh viên, cung cấp phản hồi, và đề xuất các lộ trình học tập được tùy chỉnh. Mục tiêu là tạo ra một kiến trúc linh hoạt và có khả năng mở rộng (scalable), có thể xử lý nhiều ngữ cảnh học tập khác nhau.

## Bối cảnh (Context)

Hệ thống Gia sư Thông minh (ITS) là một loại ứng dụng phần mềm sử dụng các kỹ thuật trí tuệ nhân tạo (AI) để cung cấp sự hướng dẫn cá nhân hóa cho người học. Các hệ thống này có khả năng đánh giá hiệu suất của học sinh/sinh viên, cung cấp gợi ý (hints) và đưa ra phản hồi dựa trên hành vi và tiến trình của người học. Một ITS đặt mục tiêu tái tạo một số khía cạnh của việc gia sư một kèm một (one-on-one) với người thật, trong đó hệ thống sẽ điều chỉnh linh hoạt theo tốc độ, khả năng và sở thích của người học.

Hệ thống phải có khả năng xử lý nhiều môn học, chủ đề và mức độ khó khác nhau, đồng thời cũng phải cung cấp các tính năng quản trị cho giảng viên/giáo viên, chẳng hạn như theo dõi tiến trình của học sinh/sinh viên và tạo báo cáo.

## Phạm vi của Hệ thống (Scope of the System)

Các phạm vi chính của hệ thống bao gồm:

1. **Học tập Cá nhân hóa (Personalized Learning):** Hệ thống phải đánh giá điểm mạnh, điểm yếu và phong cách học tập của từng học sinh/sinh viên, đồng thời cung cấp nội dung hướng dẫn được tùy chỉnh.

2. **Phản hồi (Feedback):** Học sinh/sinh viên phải nhận được phản hồi về hiệu suất của họ, bao gồm cả hướng dẫn và gợi ý khi cần thiết.

3. **Đánh giá và Thẩm định (Assessment and Evaluation):** ITS phải bao gồm các cơ chế để đánh giá tiến trình của học sinh/sinh viên thông qua các bài kiểm tra ngắn (quizzes), bài tập (exercises) hoặc dự án (projects).

4. **Bảng điều khiển Giảng viên/Giáo viên (Instructor Dashboard):** Giáo viên có thể giám sát tiến trình của học sinh/sinh viên, quản lý nội dung và tạo báo cáo.

5. **Quản lý Nội dung Học tập (Learning Content Management):** Khả năng quản lý, cập nhật và tuyển chọn tài liệu học tập trên nhiều chủ đề, định dạng khác nhau (ví dụ: văn bản, video, bài tập tương tác).


## Các Nhiệm vụ được Giao (Assignment Tasks)

### Nhiệm vụ 1: Thiết kế Kiến trúc Phần mềm (Software Architecture Design)

#### Yêu cầu chi tiết:
- Mô tả chi tiết bối cảnh của Hệ thống Gia sư Thông minh (ITS) dựa trên các thông tin cơ bản đã cung cấp. Điều này bao gồm việc xác định rõ ràng các yêu cầu chức năng (functional) và phi chức năng (non-functional), phác thảo các mục tiêu và phạm vi của dự án, v.v.

#### Tạo ra kiến trúc phần mềm cho ITS:
Kiến trúc phần mềm nên bao gồm:

1. **Các Đặc tính Kiến trúc (Architecture Characteristics)**: Xác định các tiêu chí thành công của ITS.

2. **Cấu trúc (Structure)**:
   - So sánh và chọn các kiểu kiến trúc (architecture styles) phù hợp để áp dụng cho ITS.
   - Trình bày kiến trúc phần mềm theo các góc nhìn (views) khác nhau, bao gồm:
     - Góc nhìn mô-đun (module views)
     - Góc nhìn thành phần và kết nối (component–and–connector views)
     - Góc nhìn phân bổ (allocation views)

3. **Các Quyết định Kiến trúc (Architecture Decisions)**: Xác định các quy tắc cho cách thức ITS nên được xây dựng.

4. **Các Nguyên tắc Thiết kế (Design Principles)**: Các hướng dẫn để xây dựng ITS.

5. **Áp dụng các Nguyên tắc SOLID (Apply SOLID Principles)**: Giải thích cách các nguyên tắc SOLID đã được áp dụng trong thiết kế của bạn.

6. **Báo cáo Phản ánh (Reflection Report)**: Viết một báo cáo phản ánh ngắn gọn về việc áp dụng các nguyên tắc SOLID đã giúp bạn cải thiện thiết kế của ITS như thế nào. Thảo luận về những thách thức bạn đã gặp phải và cách việc tuân thủ các nguyên tắc này đã làm cho hệ thống của bạn trở nên mô-đun hơn (modular), dễ bảo trì hơn (maintainable) và có khả năng mở rộng hơn (extensible).

### Nhiệm vụ 2: Triển khai Mã (Code Implementation)
*(chọn ít nhất một mô-đun chính để triển khai)*

- **Triển khai các Chức năng Cốt lõi (Implement Core Functionalities)**: Triển khai các tính năng hệ thống chính dựa trên thiết kế của bạn, đảm bảo mã nguồn tuân thủ các nguyên tắc SOLID.

## Hướng dẫn Nộp bài (Submission Instructions)

Nộp một báo cáo chi tiết có chứa đường dẫn URL GitHub (mã nguồn) trên trang web LMS.

