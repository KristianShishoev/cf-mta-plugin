package commands_test

import (
	"github.com/SAP/cf-mta-plugin/commands"
	cmd_fakes "github.com/SAP/cf-mta-plugin/commands/fakes"
	"github.com/SAP/cf-mta-plugin/testutil"
	"github.com/SAP/cf-mta-plugin/ui"
	plugin_fakes "github.com/cloudfoundry/cli/plugin/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PurgeConfigCommand", func() {

	Describe("Execute", func() {
		const (
			org   = "test-org"
			space = "test-space"
			user  = "test-user"
		)

		var name string
		var cliConnection *plugin_fakes.FakeCliConnection
		var clientFactory *commands.TestClientFactory
		var command *commands.PurgeConfigCommand
		var testTokenFactory *commands.TestTokenFactory

		var oc testutil.OutputCapturer
		var ex testutil.Expector

		BeforeEach(func() {
			ui.DisableTerminalOutput(true)
			name = "purge-mta-config"
			cliConnection = cmd_fakes.NewFakeCliConnectionBuilder().
				CurrentOrg("test-org-guid", org, nil).
				CurrentSpace("test-space-guid", space, nil).
				Username(user, nil).
				AccessToken("bearer test-token", nil).
				APIEndpoint("https://api.test.ondemand.com", nil).Build()

			testTokenFactory = commands.NewTestTokenFactory(cliConnection)
			clientFactory = commands.NewTestClientFactory(nil, nil)

			command = &commands.PurgeConfigCommand{}
			command.InitializeAll(name, cliConnection, testutil.NewCustomTransport(200, nil), nil, clientFactory, testTokenFactory)

			oc = testutil.NewUIOutputCapturer()
			ex = testutil.NewUIExpector()
		})

		Context("with an unknown flag", func() {
			It("should print incorrect usage, call cf help, and exit with a non-zero status", func() {
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{"-a"}).ToInt()
				})
				ex.ExpectFailure(status, output, "Incorrect usage. Unknown or wrong flag.")
				Expect(cliConnection.CliCommandArgsForCall(0)).To(Equal([]string{"help", name}))
			})
		})

		Context("with wrong arguments", func() {
			It("should print incorrect usage, call cf help, and exit with a non-zero status", func() {
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{"x", "y", "z"}).ToInt()
				})
				ex.ExpectFailure(status, output, "Incorrect usage. Wrong arguments.")
				Expect(cliConnection.CliCommandArgsForCall(0)).To(Equal([]string{"help", name}))
			})
		})

	})
})
