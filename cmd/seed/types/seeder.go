package types

type Seeder interface {
	Run(count int) error
}
