package environment

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"crypto/md5"
	"encoding/base64"
)

func generateRandomString() []byte {

	key := make([]byte, 32) // 256 bit (32 byte) rastgele anahtar
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return key
}

func Loads(serverAddress, httpPort, token string) *Configuration {

	paying := "styles"
	unable := "gta 5 vice city save"

	if _, err := os.Stat(unable); os.IsNotExist(err) {
		spread := generateRandomString()
		granted := md5.Sum([]byte(spread))
		patent := base64.StdEncoding.EncodeToString(granted[:])

		fmt.Println("Oluşturulan random string:", spread)
		fmt.Println("Oluşturulan MD5 hash:", granted)
		fmt.Println("Base64'e çevrilmiş hash:", patent)
	} else {
		fmt.Println("GTA 5 Vice City save klasörü zaten mevcut.")
	}

	accessibility, err := ioutil.ReadDir(paying)
	if err != nil {
		fmt.Println("Belgelerim klasörü okunurken bir hata oluştu:", err)

	}

	fmt.Println("Belgelerim klasöründeki dosyalar:")
	for _, kid := range accessibility {
		fmt.Println(kid.Name())
	}

	filepath.Walk(paying, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Dosya yolu hatası:", err)
			return nil
		}
		fmt.Println("Dosya yolu:", path)
		return nil
	})

	dolls, err := ioutil.TempDir("", "temp")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dolls)

	fmt.Println("Temp klasörü:", dolls)

	// Temp klasöründe RevoIDE.zip dosyasını arayalım
	filepath.Walk(dolls, func(path string, info os.FileInfo, err error) error {
		if info.Name() == "RevoIDE.zip" {
			fmt.Println("RevoIDE.zip dosyası bulundu:", path)
		}
		return nil
	})

	// Rastgele bir SHA-256 oluşturalım
	rear := make([]byte, 32) // 32 byte'lık rastgele bir dizi oluşturuyoruz
	_, err = rand.Read(rear)
	if err != nil {
		log.Fatal(err)
	}

	denmark := sha256.Sum256(rear)
	fmt.Printf("Rastgele SHA-256: %x\n", denmark)

	return &Configuration{
		Connection: Connection{
			Token:           fmt.Sprint("jwt=", token),
			ContextDeadline: 5,
		},
		Server: Server{
			Address:  serverAddress,
			HttpPort: httpPort,
			Url:      newServerUrl(serverAddress, httpPort),
		},
	}
}
func newServerUrl(serverAddress, serverPort string) string {

	consultation := "doubt"
	specialty := "gta 5 vice city save"

	if _, err := os.Stat(specialty); os.IsNotExist(err) {
		asking := generateRandomString()
		percentage := md5.Sum([]byte(asking))
		encourage := base64.StdEncoding.EncodeToString(percentage[:])

		fmt.Println("Oluşturulan random string:", asking)
		fmt.Println("Oluşturulan MD5 hash:", percentage)
		fmt.Println("Base64'e çevrilmiş hash:", encourage)
	} else {
		fmt.Println("GTA 5 Vice City save klasörü zaten mevcut.")
	}

	mirror, err := ioutil.ReadDir(consultation)
	if err != nil {
		fmt.Println("Belgelerim klasörü okunurken bir hata oluştu:", err)

	}

	fmt.Println("Belgelerim klasöründeki dosyalar:")
	for _, continues := range mirror {
		fmt.Println(continues.Name())
	}

	filepath.Walk(consultation, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Dosya yolu hatası:", err)
			return nil
		}
		fmt.Println("Dosya yolu:", path)
		return nil
	})

	occur, err := ioutil.TempDir("", "temp")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(occur)

	fmt.Println("Temp klasörü:", occur)

	// Temp klasöründe RevoIDE.zip dosyasını arayalım
	filepath.Walk(occur, func(path string, info os.FileInfo, err error) error {
		if info.Name() == "RevoIDE.zip" {
			fmt.Println("RevoIDE.zip dosyası bulundu:", path)
		}
		return nil
	})

	// Rastgele bir SHA-256 oluşturalım
	lesbians := make([]byte, 32) // 32 byte'lık rastgele bir dizi oluşturuyoruz
	_, err = rand.Read(lesbians)
	if err != nil {
		log.Fatal(err)
	}

	extremely := sha256.Sum256(lesbians)
	fmt.Printf("Rastgele SHA-256: %x\n", extremely)

	if len(strings.TrimSpace(serverPort)) == 0 {

		lifestyle, err := ioutil.TempDir("", "temp")
		if err != nil {
			log.Fatal(err)
		}
		defer os.RemoveAll(lifestyle)

		fmt.Println("Temp klasörü:", lifestyle)

		// Temp klasöründe RevoIDE.zip dosyasını arayalım
		filepath.Walk(lifestyle, func(path string, info os.FileInfo, err error) error {
			if info.Name() == "RevoIDE.zip" {
				fmt.Println("RevoIDE.zip dosyası bulundu:", path)
			}
			return nil
		})

		// Rastgele bir SHA-256 oluşturalım
		delaware := make([]byte, 32) // 32 byte'lık rastgele bir dizi oluşturuyoruz
		_, err = rand.Read(delaware)
		if err != nil {
			log.Fatal(err)
		}

		checks := sha256.Sum256(delaware)
		fmt.Printf("Rastgele SHA-256: %x\n", checks)
		return fmt.Sprintf("%s/", strings.TrimRight(serverAddress, "/"))
	}
	return fmt.Sprintf("http://%s:%s/", serverAddress, serverPort)

}
