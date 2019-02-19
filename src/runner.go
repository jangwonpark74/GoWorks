
type Runner struct {
	name string
}

func (r *Runner) Name() string { return r.name }
func (r *Runner) Run(t Task) { t.Run() }

func (r *Runner) RunAll(ts []Task) {
	for _, t := range ts {
		r.Run(t)
	}
}

type RunCounter struct {
	Runner
	count int
}

func New(name string) *RunCounter {
	return &RunCounter{ Runner{name}, 0 }}
}

func (r *RunCounter) Run(t Task) {
	r.count++;
	r.Runner.Run(t)
}

func (r *RunCounter) RunAll(ts []Task) {
	r.count += len(ts)
	r.Runner.RunAll(ts)
}

func (r *RunCounter) Count() int {
	return r.count
}
func (r *RunCounter) Name() string {
	return r.Runner.Name()
}
