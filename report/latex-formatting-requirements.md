# Hướng Dẫn Format LaTeX cho Báo Cáo

## Tổng Quan

Tài liệu này mô tả các quy tắc format LaTeX đã được áp dụng cho các báo cáo Problem Set 3. Các quy tắc này đảm bảo tính nhất quán, chuyên nghiệp và dễ đọc cho tài liệu LaTeX.

---

## 1. Cấu Trúc Hierarchical

### 1.1 Sections và Subsections

- Sử dụng `\section{}`, `\subsection{}`, `\subsubsection{}` theo cấu trúc phân cấp
- **Không viết hoa toàn bộ** trong tiêu đề, chỉ viết hoa chữ cái đầu của từ quan trọng

**Ví dụ:**
```latex
\section{Câu 1. Cryptographic Hardness}
\subsection{Câu 1.1. P vs NP}
\subsubsection{(a) (5 điểm) Kịch bản "Ngày Tận Thế Của Mật Mã"}
```

**Tránh:**
```latex
\section{CÂU 1. CRYPTOGRAPHIC HARDNESS}  ❌
\subsection{CÂU 1.1. P VS NP}  ❌
```

### 1.2 Format Câu Trả Lời

Mỗi phần trả lời bắt đầu với:
```latex
\begin{center}
    \textbf{Trả lời:}
\end{center}
```

Sau đó là tiêu đề phụ (nếu có):
```latex
\textbf{(a) Chỉ định các lựa chọn thuật toán mật mã}
```

---

## 2. Format Toán Học

### 2.1 Biến và Ký Hiệu

- Tất cả biến toán học phải được đặt trong `$...$` hoặc `\(...\)` cho inline math
- Sử dụng `\[...\]` hoặc `$$...$$` cho display math (không khuyến khích `$$`)

**Ví dụ:**
```latex
Khóa bí mật chung $S$ được tính từ $S = ab \cdot G$.
Công thức: $d = r^{-1} \cdot (s \cdot k - h) \pmod{n}$
```

### 2.2 Các Phép Toán

- Phép nhân: `$a \cdot b$` hoặc `$a \times b$` (không dùng `*` trong math mode)
- Phép lũy thừa: `$2^{256}$`, `$k^{-1}$`
- Phép modulo: `$a \pmod{n}$` hoặc `$a \bmod n$`
- Phép cộng/trừ: `$a + b$`, `$a - b$`
- Phép so sánh: `$S < L$`, `$k \leq n$`
- Mũi tên: `$\rightarrow$` hoặc `$\to$` cho `→`

**Ví dụ:**
```latex
$A = a \cdot G$ và $B = b \cdot G$
$S = ab \cdot G$
$s = k^{-1} \cdot (H(m) + r \cdot d) \pmod{n}$
```

### 2.3 Ký Hiệu Đặc Biệt

- Tập hợp số nguyên: `$\mathbb{Z}$`, `$\mathbb{Z}_p^*$`
- Xấp xỉ: `$\sim$` cho `~`
- Dấu ngoặc kép trong math: `$``...''$` (không dùng `"`)

---

## 3. Format Code Blocks

### 3.1 Verbatim Environment

Sử dụng `\begin{verbatim}...\end{verbatim}` cho code blocks:

```latex
\begin{verbatim}
def scalar_mult_naive(k, G):
    """
    Tinh k*G bang double-and-add
    KHONG AN TOAN - co timing leak
    """
    result = POINT_AT_INFINITY
    temp = G
    
    for bit in bits(k):
        if bit == 1:
            result = result + temp
        temp = temp + temp
    
    return result
\end{verbatim}
```

**Lưu ý:**
- Loại bỏ emoji và special characters trong code
- Chuyển Unicode sang ASCII khi cần (ví dụ: `→` thành `->`)
- Không dùng box-drawing characters Unicode (như `┌`, `─`, `│`), thay bằng ASCII (`+`, `-`, `|`)

### 3.2 Inline Code

Sử dụng `\texttt{}` cho tên file, biến, hoặc code ngắn:

```latex
File \texttt{main.tex} chứa cấu hình chính.
```

---

## 4. Format Lists

### 4.1 Itemize (Danh sách không đánh số)

```latex
\begin{itemize}
    \item Mục đầu tiên
    \item Mục thứ hai
    \item Mục thứ ba
\end{itemize}
```

### 4.2 Enumerate (Danh sách có đánh số)

```latex
\begin{enumerate}
    \item Bước đầu tiên
    \item Bước thứ hai
    \item Bước thứ ba
\end{enumerate}
```

### 4.3 Nested Lists

```latex
\begin{itemize}
    \item \textbf{Tiêu đề phụ:}
    \begin{enumerate}
        \item Chi tiết 1
        \item Chi tiết 2
    \end{enumerate}
    \item \textbf{Tiêu đề phụ khác:}
    \begin{itemize}
        \item Chi tiết a
        \item Chi tiết b
    \end{itemize}
\end{itemize}
```

---

## 5. Format Tables

### 5.1 Basic Table

```latex
\begin{center}
\begin{tabular}{|l|l|l|}
\hline
\textbf{Cột 1} & \textbf{Cột 2} & \textbf{Cột 3} \\
\hline
Dòng 1, Cột 1 & Dòng 1, Cột 2 & Dòng 1, Cột 3 \\
\hline
Dòng 2, Cột 1 & Dòng 2, Cột 2 & Dòng 2, Cột 3 \\
\hline
\end{tabular}
\end{center}
```

### 5.2 Table với Math

```latex
\begin{center}
\begin{tabular}{|l|l|l|}
\hline
\textbf{Kỹ thuật} & \textbf{Overhead} & \textbf{Độ an toàn} \\
\hline
Naive implementation & $0\%$ & Không an toàn \\
\hline
Montgomery Ladder & $+15\%$ & An toàn \\
\hline
\end{tabular}
\end{center}
```

---

## 6. Format Text Styling

### 6.1 Bold và Italic

- **Bold:** `\textbf{text}` cho tiêu đề, từ khóa quan trọng
- **Italic:** `\textit{text}` cho nguồn tham khảo, thuật ngữ

**Ví dụ:**
```latex
\textbf{(a) Chỉ định các lựa chọn thuật toán mật mã}
\textit{Nguồn:} Mục 2.3 (Chương 8) - ``EdDSA: Tiêu chuẩn hiện đại...''
```

### 6.2 Quotes

- Sử dụng dấu ngoặc kép LaTeX: `` `...' `` cho quotes đơn, `` ``...'' `` cho quotes kép

**Ví dụ:**
```latex
EdDSA là ``xác định'' (deterministic).
Tài liệu khuyến nghị ``an toàn và nhanh hơn''.
```

---

## 7. Xử Lý Special Characters

### 7.1 Emoji và Symbols

**Loại bỏ hoặc thay thế:**
- ✅ → `\textbf{✓}` hoặc bỏ
- ❌ → `\textbf{✗}` hoặc bỏ
- ⚡ → bỏ hoặc thay bằng text
- 📍 → `\textit{Nguồn:}` hoặc bỏ
- 🔴, 🟠, 🟡, 🟢 → bỏ hoặc thay bằng text mô tả

**Ví dụ:**
```latex
❌ Không an toàn  →  \textbf{Không} an toàn
✅ An toàn  →  An toàn
📍 Nguồn: ...  →  \textit{Nguồn:} ...
```

### 7.2 Unicode Characters

- Chuyển Unicode sang LaTeX equivalents:
  - `√` → `$\sqrt{}$`
  - `Σ` → `$\sum$`
  - `·` → `$\cdot$`
  - `→` → `$\rightarrow$` hoặc `$\to$`
  - `≤`, `≥` → `$\leq$`, `$\geq$`
  - `≠` → `$\neq$`
  - `≈` → `$\approx$`

### 7.3 Escape Characters

- `&` → `\&` (trong text mode)
- `%` → `\%`
- `$` → `\$` (trong text mode)
- `#` → `\#`
- `_` → `\_` (trong text mode)
- `^` → `\^{}` (trong text mode)
- `{` → `\{`
- `}` → `\}`

---

## 8. Format Nguồn Tham Khảo

### 8.1 Inline Citations

```latex
\textit{Nguồn:} Mục 2.3 (Chương 8) - ``EdDSA: Tiêu chuẩn hiện đại (2011)...''
```

### 8.2 Source Lists

```latex
\textbf{Nguồn tham chiếu:}
\begin{itemize}
    \item Mục 2.3 (Chương 8): ECDSA và EdDSA
    \item Mục 3.2 (Chương 8): Curve25519 an toàn hơn
    \item Ví dụ 2 (Chương 8): PlayStation 3 hack - bài học về ECDSA
\end{itemize}
```

---

## 9. Spacing và Layout

### 9.1 Vertical Spacing

Sử dụng `\vspace{}` khi cần thiết:
```latex
\vspace{0.5cm}
```

### 9.2 No Indent

Sử dụng `\noindent` khi cần bắt đầu đoạn không thụt lề:
```latex
\noindent\textbf{Tiêu đề:} Nội dung...
```

### 9.3 Line Breaks

- Sử dụng `\\` cho line break trong tables
- Sử dụng blank line (`\n\n`) để tạo paragraph break

---

## 10. Quy Tắc Viết Hoa

### 10.1 Tiêu Đề

- **Không viết hoa toàn bộ** trừ khi là từ viết tắt (VD: ECDSA, MITM, CA)
- Chỉ viết hoa chữ cái đầu của từ quan trọng

**Ví dụ:**
```latex
✅ \subsection{Câu 2. Kiến trúc hệ thống liên lạc an toàn}
❌ \subsection{CÂU 2. KIẾN TRÚC HỆ THỐNG LIÊN LẠC AN TOÀN}
```

### 10.2 Tiêu Đề Phụ

```latex
✅ \textbf{(a) Chỉ định các lựa chọn thuật toán mật mã}
✅ \textbf{1. Khởi tạo lòng tin (Bootstrap Trust):}
✅ \textbf{2. Phân phối khóa (Key Distribution):}
```

---

## 11. Cấu Trúc Câu Trả Lời Hoàn Chỉnh

### 11.1 Template Cơ Bản

```latex
\begin{center}
    \textbf{Trả lời:}
\end{center}

\textbf{(a) Tiêu đề câu trả lời}

Nội dung giới thiệu...

\textbf{1. Tiêu đề phụ đầu tiên:}

\begin{itemize}
    \item Chi tiết 1
    \item Chi tiết 2
\end{itemize}

\textbf{2. Tiêu đề phụ thứ hai:}

\begin{enumerate}
    \item Bước 1
    \item Bước 2
\end{enumerate}

\textit{Nguồn:} Mục X.X (Chương Y) - ``Mô tả nguồn...''
```

### 11.2 Template với Code

```latex
\begin{center}
    \textbf{Trả lời:}
\end{center}

\textbf{(a) Thiết kế tấn công timing}

\textbf{A. Kiến thức nền}

\textit{Nguồn:} Mục 2.3 (Chương 8) - ``ECDSA: Tiêu chuẩn cũ...''

\textbf{Quy trình ký ECDSA:}
\begin{enumerate}
    \item Chọn nonce ngẫu nhiên: $k$
    \item Tính điểm: $R = k \cdot G$
\end{enumerate}

\textbf{B. Điểm yếu có thể khai thác}

\begin{verbatim}
def attack_function():
    # Code here
    pass
\end{verbatim}
```

---

## 12. Checklist Format

Trước khi hoàn thành, kiểm tra:

- [ ] Tất cả tiêu đề không viết hoa toàn bộ
- [ ] Tất cả biến toán học trong `$...$`
- [ ] Code blocks trong `\begin{verbatim}...\end{verbatim}`
- [ ] Loại bỏ hoặc thay thế emoji
- [ ] Sử dụng dấu ngoặc kép LaTeX (``...'')
- [ ] Escape special characters đúng cách
- [ ] Mỗi phần trả lời có `\begin{center}\textbf{Trả lời:}\end{center}`
- [ ] Tables được đặt trong `\begin{center}...\end{center}`
- [ ] Lists sử dụng `\begin{itemize}` hoặc `\begin{enumerate}`
- [ ] Nguồn tham khảo format với `\textit{Nguồn:}`

---

## 13. Ví Dụ Hoàn Chỉnh

```latex
\subsubsection{(a) (8 điểm) Phòng thí nghiệm tấn công kênh bên}

Bạn được giao nhiệm vụ kiểm thử triển khai ECDSA...

\begin{center}
    \textbf{Trả lời:}
\end{center}

\textbf{i. Thiết kế tấn công timing chống ECDSA}

\textbf{A. Kiến thức nền về ECDSA}

\textit{Nguồn:} Mục 2.3 (Chương 8) - ``ECDSA: Tiêu chuẩn cũ, được NIST chuẩn hóa. \textbf{Điểm yếu chí mạng:} Khi ký, ECDSA yêu cầu một số ngẫu nhiên bí mật $k$ (gọi là 'nonce'). Nếu $k$ bị rò rỉ, bị lặp lại, hoặc có thể dự đoán, khóa riêng (private key) $d$ sẽ bị lộ ngay lập tức!''

\textbf{Quy trình ký ECDSA (đơn giản hóa):}
\begin{enumerate}
    \item Chọn nonce ngẫu nhiên: $k$
    \item Tính điểm: $R = k \cdot G$
    \item Lấy tọa độ $x$: $r = R.x \pmod{n}$
    \item Tính: $s = k^{-1} \cdot (H(m) + r \cdot d) \pmod{n}$
    \item Chữ ký: $(r, s)$
\end{enumerate}

\textbf{B. Điểm yếu có thể khai thác}

\textbf{Các phép toán có thời gian khác nhau:}

\textbf{1. Scalar multiplication: $k \cdot G$}
\begin{itemize}
    \item Phụ thuộc vào các bit của $k$
    \item Có thể phân biệt bit 0 vs bit 1
\end{itemize}

\textbf{2. Modular inversion: $k^{-1}$}
\begin{itemize}
    \item Có thể phụ thuộc vào giá trị $k$
    \item Một số thuật toán không constant-time
\end{itemize}
```

---

## 14. Lưu Ý Quan Trọng

1. **Không thêm nội dung:** Chỉ format, không thêm hoặc xóa nội dung
2. **Giữ nguyên ý nghĩa:** Đảm bảo format không làm thay đổi ý nghĩa
3. **Nhất quán:** Áp dụng cùng một style cho toàn bộ document
4. **Kiểm tra lỗi:** Sau khi format, chạy `read_lints` để kiểm tra lỗi LaTeX
5. **Tối ưu:** Ưu tiên readability và professional appearance

---

## 15. Các Lỗi Thường Gặp và Cách Sửa

### 15.1 Unicode trong Verbatim

**Lỗi:**
```latex
\begin{verbatim}
┌─────────────────────┐
│   Unicode box       │
└─────────────────────┘
\end{verbatim}
```

**Sửa:**
```latex
\begin{verbatim}
+---------------------+
|   ASCII box         |
+---------------------+
\end{verbatim}
```

### 15.2 Math Mode Thiếu

**Lỗi:**
```latex
Giá trị k nhỏ hơn n
```

**Sửa:**
```latex
Giá trị $k$ nhỏ hơn $n$
```

### 15.3 Quotes Sai

**Lỗi:**
```latex
EdDSA là "xác định"
```

**Sửa:**
```latex
EdDSA là ``xác định''
```

---

## Kết Luận

Tài liệu này cung cấp hướng dẫn chi tiết để format LaTeX một cách nhất quán và chuyên nghiệp. Khi format, hãy tham khảo các quy tắc trên và sử dụng các template có sẵn để đảm bảo tính nhất quán.

