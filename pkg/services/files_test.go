package services_test

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/redhatinsights/edge-api/config"
	"github.com/redhatinsights/edge-api/pkg/services"
	log "github.com/sirupsen/logrus"
)

var _ = Describe("File Service Test", func() {
	var logEntry *log.Entry
	Describe("local file service", func() {
		var service services.FilesService

		BeforeEach(func() {
			logEntry = log.NewEntry(log.StandardLogger())
			cfg := config.Get()
			cfg.Local = true
			service = services.NewFilesService(logEntry)
		})
		When("file service is created", func() {
			It("return service", func() {
				Expect(service).To(Not(BeNil()))
			})
		})
		When("get file", func() {
			var path, filename, data string
			BeforeEach(func() {
				data = "i am a file data"
				filename = "test"
				path = fmt.Sprintf("/tmp/%s", filename)
				f, err := os.Create(path)
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()
				ioutil.WriteFile(path, []byte(data), fs.ModeAppend)
			})
			AfterEach(func() {
				os.Remove(path)
			})

			It("returns file", func() {
				file, err := service.GetFile(filename)
				Expect(err).To(BeNil())

				b, err := io.ReadAll(file)
				Expect(err).To(BeNil())
				Expect(string(b)).To(Equal(data))
			})
		})
	})
	Describe("aws file service", func() {
		BeforeEach(func() {
			logEntry = log.NewEntry(log.StandardLogger())
			cfg := config.Get()
			cfg.Local = false
			cfg.Debug = true
		})
		When("aws file service is created", func() {
			var service services.FilesService
			BeforeEach(func() {
				service = services.NewFilesService(logEntry)
			})
			It("return service", func() {
				Expect(service).To(Not(BeNil()))
			})
		})
	})
})
