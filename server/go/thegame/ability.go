package main

const (
	HealthRegen = iota
	MaxHealth
	BodyDamage
	BulletSpeed
	BulletPenetration
	BulletDamage
	Reload
	MovementSpeed
	NAbilities
)

func linearScale(at0, at8 int) [9]int {
	step8 := at8 - at0
	L := func(i int) int {
		return at0 + int(float64((step8)*i)/8)
	}
	return [9]int{L(0), L(1), L(2), L(3), L(4), L(5), L(6), L(7), L(8)}
}

var AbilityValues = [NAbilities][9]int{
	{0, 1, 1, 2, 2, 3, 3, 7, 8},                                   // HealthRegen
	{1000, 1500, 2000, 2500, 3000, 3500, 4000, 5200, 6000},        // MaxHealth
	{15, 18, 22, 27, 31, 35, 40, 50, 58},                          // BodyDamage
	linearScale(8, 24),                                            // BulletSpeed
	linearScale(40, 136),                                          // BulletPenetration
	{18, 19, 20, 21, 28, 35, 42, 52, 62},                          // BulletDamage
	{15, 14, 12, 9, 7, 6, 5, 4, 2},                                // Reload
	linearScale(6, 14),                                            // MovementSpeed
}
