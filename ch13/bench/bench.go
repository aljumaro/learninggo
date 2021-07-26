package bench

import "os"

func FileLen(f string, bufferSize int) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	count := 0
	buffer := make([]byte, bufferSize)
	for {
		num, err := file.Read(buffer)
		count += num
		if err != nil {
			break
		}
	}
	return count, nil
}
