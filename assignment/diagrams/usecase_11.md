
<span style="font-size: 24px;"><b>UC-11 – Gợi ý Bài học Bù (Remediation)</b></span>

<span style="font-size: 18px;">
<b>Mục tiêu:</b> Khi học viên yếu kỹ năng nào đó, hệ thống đề xuất nội dung phù hợp để củng cố.
</span>

```mermaid
graph TD
    subgraph ITS
        UC10["UC-10: Chấm điểm & Phản hồi"]
        UC11["UC-11: Gợi ý Bài học Bù"]
        FR6["Feedback & Remediation Service"]
        FR7["Adaptive Learning Engine"]
    end
    Learner
    Instructor
    UC10 -- "extends" --> UC11
    UC11 -- "includes" --> FR6
    UC11 -- "includes" --> FR7
    Learner --> UC11
    Instructor -. review .-> UC11
```
<span style="font-size: 18px;">
- Kích hoạt khi người học sai liên tục hoặc mastery score &lt; 0.6.<br>
- Adaptive Learning Engine và Feedback Service cùng xử lý để sinh danh sách bài bù.<br>
- Instructor có thể xem và xác nhận lộ trình ôn tập.
</span>

<span style="font-size: 20px;"><b>Tổng quan chuỗi Use Case 09 → 10 → 11</b></span>

```mermaid
graph LR
    UC09["UC-09: Làm Bài tập"] --> UC10["UC-10: Chấm điểm & Phản hồi"]
    UC10 --> UC11["UC-11: Gợi ý Bài học Bù"]
```