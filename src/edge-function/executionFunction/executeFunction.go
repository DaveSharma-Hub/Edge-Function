package executeFunction

import (
	"os/exec"
	"strings"
	"fmt"
	"github.com/google/uuid"
)

func runCommand(command string) (string,error) {
	commands := strings.Split(command, " ");
	cmd := exec.Command(commands[0], commands[1:]...);
	output, err := cmd.Output();

	if err != nil {
		return "", err;
	} 

	return string(output), nil;
}

func tagDockerImage(dockerImageUri string)(string, error){
	id, err := uuid.NewRandom();
	if (err != nil){
		return "", err;
	}
	command := fmt.Sprintf("docker tag %s %s", dockerImageUri, id.String());
	_, tagError := runCommand(command);
	if(tagError != nil){
		return "", tagError;
	} 
	return id.String(), nil;
}

func executeRunDockerContainer(tagId string, timeout int, cpu float32, memory int)(string, error){
	// command := fmt.Sprintf("docker run --timeout=%d %s", timeout, tagId);
	command := fmt.Sprintf("docker run --rm --timeout=%d --memory=%dm --cpus=%d %s", timeout, memory, cpu, tagId);
	buildOutput, buildErr := runCommand(command);
	return buildOutput, buildErr;
}


func ExecuteFunction(dockerImageUri string, timeout int, cpu float32, memory int)(string, error){
	id, err := tagDockerImage(dockerImageUri);	
	if(err != nil){
		return "", err;
	}
	output, buildErr := executeRunDockerContainer(id, timeout, cpu, memory);	
	if(buildErr != nil){
		return "", buildErr;
	}
	return output, nil;
}