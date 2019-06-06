package utils

import (
	"os/exec"
	"time"

	"github.com/kubeedge/kubeedge/tests/e2e/constants"
)

//GenerateCerts - Generates Cerificates for Edge and Cloud nodes copy to respective folders
func GenerateCerts() error {
	cmd := exec.Command("bash", "-x", "scripts/generate_cert.sh")
	if err := PrintCombinedOutput(cmd); err != nil {
		return err
	}
	return nil
}

func StartEdgeController() error {
	//Run ./edgecontroller binary
	cmd := exec.Command("sh", "-c", constants.RunController)
	if err := PrintCombinedOutput(cmd); err != nil {
		return err
	}
	//Expect(err).Should(BeNil())
	time.Sleep(5 * time.Second)
	return nil
}

func StartEdgeCore() error {
	//Run ./edge_core after node registration
	cmd := exec.Command("sh", "-c", constants.RunEdgecore)
	if err := PrintCombinedOutput(cmd); err != nil {
		return err
	}
	//Expect(err).Should(BeNil())
	time.Sleep(5 * time.Second)
	return nil
}

func DeploySetup(ctx *TestContext, nodeName, setupType string) error {
	//Do the neccessary config changes in Cloud and Edge nodes
	cmd := exec.Command("bash", "-x", "scripts/setup.sh", setupType, nodeName, ctx.Cfg.K8SMasterForKubeEdge)
	if err := PrintCombinedOutput(cmd); err != nil {
		return err
	}
	//Expect(err).Should(BeNil())
	time.Sleep(1 * time.Second)
	return nil
}

func CleanUp(setupType string) error {
	cmd := exec.Command("bash", "-x", "scripts/cleanup.sh", setupType)
	if err := PrintCombinedOutput(cmd); err != nil {
		return err
	}
	time.Sleep(2 * time.Second)
	return nil
}

func StartEdgeSite() error {
	//Run ./edge_core after node registration
	cmd := exec.Command("sh", "-c", constants.RunEdgeSite)
	if err := PrintCombinedOutput(cmd); err != nil {
		return err
	}
	//Expect(err).Should(BeNil())
	time.Sleep(5 * time.Second)
	return nil
}
