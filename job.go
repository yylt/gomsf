package gomsf

import "strconv"

const (
	jobinfo = "job.info"
	joblist = "job.list"
	jobstop = "job.stop"
)

type Job struct {
	Id   int
	Name string
}

type JobDetail struct {
	Id        int                    `msgpack:"jid"`
	Name      string                 `msgpack:"name"`
	StartT    int64                  `msgpack:"start_time"`
	Uri       string                 `msgpack:"uripath"`
	DataStore map[string]interface{} `msgpack:"datastore"`
}

func JobList(cli MsfCli, token string) ([]*Job, error) {
	var (
		sts   = []string{joblist, token}
		infos []*Job
	)
	var retvar map[string]string
	err := cli.Send(sts, &retvar)
	if err != nil {
		return nil, err
	}
	for k, v := range retvar {
		kint, err := strconv.Atoi(k)
		if err != nil {
			continue
		}
		infos = append(infos, &Job{
			Id:   kint,
			Name: v,
		})
	}
	return infos, nil
}

func JobStop(cli MsfCli, token, jobid string) (*Generic, error) {
	var (
		retv *Generic
		sts  = []string{jobstop, token, jobid}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}
func JobInfo(cli MsfCli, token, jobid string) (*JobDetail, error) {
	var (
		retv *JobDetail
		sts  = []string{jobinfo, token, jobid}
	)
	err := cli.Send(sts, &retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}
