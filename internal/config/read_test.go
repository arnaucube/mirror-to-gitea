package config_test

import (
	"github.com/jaedle/mirror-to-gitea/internal/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

const mirrorMode = "MIRROR_MODE"
const giteaUrl = "GITEA_URL"
const giteaToken = "GITEA_TOKEN"

const publicMirrorMode = "PUBLIC"
const privateAndPublicMirrorMode = "PRIVATE_AND_PUBLIC"
const unknownMirrorMode = "UNKNOWN"

var _ = Describe("Read", func() {
	BeforeEach(func() {
		os.Clearenv()
	})

	It("parses valid configuration", func() {
		aValidEnv()

		c, err := config.Read()
		Expect(err).NotTo(HaveOccurred())
		Expect(c).ToNot(BeNil())
	})

	Context("Gitea", func() {
		It("parses configuration", func() {
			aValidEnv()
			setEnv(giteaUrl, "https://gitea.url/api")
			setEnv(giteaToken, "a-gitea-token")

			c, err := config.Read()

			Expect(err).NotTo(HaveOccurred())
			Expect(c.Gitea.GiteaUrl).To(Equal("https://gitea.url/api"))
			Expect(c.Gitea.GiteaToken).To(Equal("a-gitea-token"))
		})

		It("fails on missing gitea url", func() {
			aValidEnv()
			unsetEnv(giteaUrl)

			c, err := config.Read()

			Expect(err).To(HaveOccurred())
			Expect(c).To(BeNil())
		})

		It("fails on missing gitea token", func() {
			aValidEnv()
			unsetEnv(giteaToken)

			c, err := config.Read()

			Expect(err).To(HaveOccurred())
			Expect(c).To(BeNil())
		})
	})

	Context("mirror mode", func() {
		It("sets default mirror mode", func() {
			aValidEnv()
			unsetEnv(mirrorMode)

			c, err := config.Read()

			Expect(err).NotTo(HaveOccurred())
			Expect(c.MirrorMode).To(Equal(config.MirrorModePublic))
		})

		It("allows public mirror mode PUBLIC", func() {
			aValidEnv()
			setEnv(mirrorMode, publicMirrorMode)

			c, err := config.Read()

			Expect(err).NotTo(HaveOccurred())
			Expect(c.MirrorMode).To(Equal(config.MirrorModePublic))
		})

		It("allows mirror mode PRIVATE_AND_PUBLIC", func() {
			aValidEnv()
			setEnv(mirrorMode, privateAndPublicMirrorMode)

			c, err := config.Read()

			Expect(err).NotTo(HaveOccurred())
			Expect(c.MirrorMode).To(Equal(config.MirrorModePrivateAndPublic))
		})

		It("fails on unknown mirror mode", func() {
			aValidEnv()
			setEnv(mirrorMode, unknownMirrorMode)

			c, err := config.Read()

			Expect(err).To(HaveOccurred())
			Expect(c).To(BeNil())
		})

	})
})

func setEnv(k string, v string) {
	err := os.Setenv(k, v)
	Expect(err).NotTo(HaveOccurred())
}

func unsetEnv(k string) {
	err := os.Unsetenv(k)
	Expect(err).NotTo(HaveOccurred())
}

func aValidEnv() {
	setEnv(mirrorMode, "PUBLIC")
	setEnv(giteaUrl, "https://gitea.url")
	setEnv(giteaToken, "valid")
}
