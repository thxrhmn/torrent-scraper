# Menggunakan golang sebagai base image
FROM golang:latest

# Set working directory di dalam container
WORKDIR /app

# Menyalin semua file Go yang diperlukan ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o main .

# Menjalankan aplikasi saat container dijalankan
CMD ["./main"]