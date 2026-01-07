package linux

import (
	"bootstrap-open-academia/utils"
)

func BuildInstaller() {
	utils.Run("cd ../server && bun install")
	utils.Run("bun run dev")
}
