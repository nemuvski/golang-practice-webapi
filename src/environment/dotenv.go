package environment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

/*
	.env ファイルを読み込み、環境変数にセットする。
	ファイルが存在しない場合は、読み取りをスキップする。
*/
func LoadEnvFile(filePath string) {
	// ファイルがある場合のみ読み込む
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		// 事前にファイルの有無は確認しているので、エラーは無視する
		godotenv.Load(filePath)
		fmt.Printf("%s ファイルの内容を読み取りました\n", filePath)
	} else {
		fmt.Printf("%s ファイルが見つからないため、読み取りはスキップしました\n", filePath)
	}
}
