package code

import "code/internal/app"

func GenDiff(filepath1, filepath2, format string) (string, error) {
	return app.Run(filepath1, filepath2, format)
}
