package api

import (
	"dd-server/types"
	"github.com/spf13/cast"
)

// Processes

//root@bogon:/# ps -aux
//USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
//root         1  0.1  0.5  90012  2660 ?        Ss   08:36   0:01 /opt/datadog-ag
//root        10  0.3  1.5 204900  8008 ?        Sl   08:36   0:03 /opt/datadog-ag
//root        12  0.5  1.7 142040  8660 ?        S    08:36   0:06 /opt/datadog-ag
//root        13  0.3  2.4 139040 12408 ?        S    08:36   0:03 /opt/datadog-ag
//root        33  0.0  0.1  20228   848 ?        Ss   08:37   0:00 bash
//root       380  0.0  0.2  17492  1136 ?        R+   08:56   0:00 ps -aux

func ParseProcesses(ps *types.ProcessStruct) []types.Process {

	processes := make([]types.Process, 0)

	for _, process := range ps.Processes {
		if pa, ok := process.([]interface{}); ok {
			if len(pa) != 11 {
				continue
			}
			p := types.Process{
				User:    cast.ToString(pa[0]),
				Process: cast.ToString(pa[10]),
			}
			processes = append(processes, p)
		}

	}

	return processes
}
