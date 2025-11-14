<span style="font-size: 24px;"><b>UC-09 – Làm Bài tập và Assessment</b></span>

```mermaid
graph TD
    subgraph ITS
        UC05["UC-05: Tạo Bài tập"]
        UC09["UC-09: Làm Bài tập & Assessment"]
        UC10["UC-10: Chấm điểm & Phản hồi Tức thì"]
    end
    Learner((Learner Học sinh))
    Instructor((Instructor Giảng viên))
    Instructor --> UC05
    UC05 -- "includes" --> UC09
    Learner --> UC09
    UC09 -- "extends" --> UC10
```

<span style="font-size: 18px;">
<b>Giải thích:</b><br>
- Instructor tạo nội dung (UC-05).<br>
- Learner thực hiện bài tập (UC-09).<br>
- Khi nộp bài, hệ thống mở rộng (extends) đến UC-10 để chấm điểm và phản hồi.
</span>