package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Đang bắt đầu sinh tài liệu Swagger... Vui lòng đợi trong giây lát...")
	
	// Chạy lệnh swag init thông qua go run
	cmd := exec.Command("go", "run", "github.com/swaggo/swag/cmd/swag@latest", "init", "-g", "cmd/api/main.go")
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Lỗi khi sinh Swagger: %v\n", err)
		fmt.Printf("Chi tiết lỗi:\n%s\n", string(output))
		fmt.Println("\nNếu bạn thấy lỗi 'go: command not found', hãy thử khởi động lại IDE.")
		return
	}

	fmt.Println("==================================================")
	fmt.Println("✅ CHÚC MỪNG! ĐÃ SINH TÀI LIỆU SWAGGER THÀNH CÔNG.")
	fmt.Println("Bây giờ bạn hãy chạy lại server và tận hưởng kết quả nhé.")
	fmt.Println("==================================================")
}
