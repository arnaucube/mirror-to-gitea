package config_test

import (
	"fmt"
	"github.com/jaedle/mirror-to-gitea/internal/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"os"
)

const mirrorMode = "MIRROR_MODE"
const giteaUrl = "GITEA_URL"
const giteaToken = "GITEA_TOKEN"
const githubUsername = "GITHUB_USERNAME"
const githubToken = "GITHUB_TOKEN"

const publicMirrorMode = "PUBLIC"
const privateAndPublicMirrorMode = "PRIVATE_AND_PUBLIC"
const unknownMirrorMode = "UNKNOWN"

const aGithubUsername = "a-github-user"
const aGithubToken = "a-github-token"

var _ = Describe("Read", func() {

	var reader *config.Reader

	BeforeEach(func() {
		os.Clearenv()
		reader = config.NewReader()
	})

	It("parses valid configuration", func() {
		aValidEnv()

		c, err := reader.Read()
		Expect(err).NotTo(HaveOccurred())
		Expect(c).ToNot(BeNil())
	})

	Context("github", func() {
		It("parses configuration", func() {
			aValidEnv()
			setEnv(githubUsername, aGithubUsername)
			unsetEnv(githubToken)

			c, err := reader.Read()

			Expect(err).NotTo(HaveOccurred())
			Expect(c.Github.Username).To(Equal(aGithubUsername))
			Expect(c.Github.Token).To(BeNil())
		})

		It("parses configuration with token", func() {
			aValidEnv()
			setEnv(githubUsername, aGithubUsername)
			setEnv(githubToken, aGithubToken)

			c, err := reader.Read()

			Expect(err).NotTo(HaveOccurred())
			Expect(c.Github.Username).To(Equal(aGithubUsername))
			Expect(*c.Github.Token).To(Equal(aGithubToken))
		})

		It("fails on missing username", func() {
			aValidEnv()
			unsetEnv(githubUsername)

			c, err := reader.Read()

			Expect(err).To(HaveOccurred())
			Expect(c).To(BeNil())
		})

	})

	Context("Gitea", func() {
		It("parses configuration", func() {
			aValidEnv()
			setEnv(giteaUrl, "https://gitea.url/api")
			setEnv(giteaToken, "a-gitea-token")

			c, err := reader.Read()

			Expect(err).NotTo(HaveOccurred())
			Expect(c.Gitea.Url).To(Equal("https://gitea.url/api"))
			Expect(c.Gitea.Token).To(Equal("a-gitea-token"))
		})

		It("fails on missing gitea url", func() {
			aValidEnv()
			unsetEnv(giteaUrl)

			c, err := reader.Read()

			Expect(err).To(HaveOccurred())
			Expect(c).To(BeNil())
			Expect(err.Error()).To(Equal("missing mandatory parameter GITEA_URL, please specify your target gitea instance"))
		})

		It("fails on missing gitea token", func() {
			aValidEnv()
			unsetEnv(giteaToken)

			c, err := reader.Read()

			Expect(err).To(HaveOccurred())
			Expect(c).To(BeNil())
			Expect(err.Error()).To(Equal("missing mandatory parameter GITEA_TOKEN, please specify your gitea application token"))

		})
	})

	Context("mirror mode", func() {
		It("sets default mirror mode", func() {
			aValidEnv()
			unsetEnv(mirrorMode)

			c, err := reader.Read()

			Expect(err).NotTo(HaveOccurred())
			Expect(c.MirrorMode).To(Equal(config.MirrorModePublic))
		})

		DescribeTable("parses mirror mode: ", func(in string, exp string) {
			aValidEnv()
			setEnv(mirrorMode, in)

			c, err := reader.Read()

			Expect(err).NotTo(HaveOccurred())
			Expect(c.MirrorMode).To(Equal(exp))
		},
			Entry("public mirror mode", publicMirrorMode, config.MirrorModePublic),
			Entry("private mirror mode", privateAndPublicMirrorMode, config.MirrorModePrivateAndPublic),
		)

		It("fails on unknown mirror mode", func() {
			aValidEnv()
			setEnv(mirrorMode, unknownMirrorMode)

			c, err := reader.Read()

			Expect(err).To(HaveOccurred())
			Expect(c).To(BeNil())

			expected := "unknown mirror mode %s, please specify a valid mirror mode: PUBLIC, PRIVATE_AND_PUBLIC"
			Expect(err.Error()).To(Equal(fmt.Sprintf(expected, unknownMirrorMode)))
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
	setEnv(githubUsername, "a-github-username")
}
