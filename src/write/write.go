package write
import "fmt"
import "io"

func WriteTo(w io.Writer, lines []string) error {
	for _, line := range lines {
		if _, err = fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return nil
}
