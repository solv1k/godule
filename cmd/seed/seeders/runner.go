package seeders

type Runner struct {
	Registrator *Registrator
}

func NewRunner() *Runner {
	return &Runner{
		Registrator: NewRegistrator(),
	}
}

// Runing database seeder using key with specified count
func (r *Runner) Run(key string, count int) error {
	seeder, err := r.Registrator.Get(key)
	if err != nil {
		return err
	}

	seeder.Run(count)
	return nil
}
