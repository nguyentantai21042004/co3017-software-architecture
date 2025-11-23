
<span style="font-size: 24px;"><b>UC-10 – Chấm điểm và Phản hồi Tức thì</b></span>

<span style="font-size: 18px;">
<b>Mục tiêu:</b> Tự động chấm điểm và trả phản hồi nhanh cho người học (&lt; 1 s), đồng thời cập nhật Learner Model.
</span>

```mermaid
graph TD
    subgraph ITS
        UC09["UC-09: Làm Bài tập"]
        UC10["UC-10: Chấm điểm & Phản hồi Tức thì"]
        FR5["Auto-Grading Service"]
        FR6["Feedback Generator"]
        UC11["UC-11: Gợi ý Bài học Bù"]
    end
    Learner
    Instructor
    Learner --> UC09
    UC09 -- "extends" --> UC10
    UC10 -- "includes" --> FR5
    UC10 -- "includes" --> FR6
    UC10 -- "extends" --> UC11
    Instructor -. monitor .-> UC10
```

<span style="font-size: 18px;">
<b>Giải thích:</b><br>
- UC-10 bao gồm hai chức năng chính:<br>
  - <b>Auto-Grading Service</b>: tính điểm.<br>
  - <b>Feedback Generator</b>: tạo gợi ý/hints.<br>
- Nếu kết quả kém, hệ thống mở rộng đến UC-11 để đề xuất bài học bù.
</span>