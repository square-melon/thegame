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
	{0, 1, 1, 2, 2, 3, 3, 4, 5},                                   // HealthRegen
	{1400, 1800, 2200, 2700, 3200, 3700, 4300, 5500, 6500},        // MaxHealth
	{15, 18, 22, 27, 33, 38, 45, 58, 68},                          // BodyDamage
	linearScale(8, 20),                                            // BulletSpeed
	linearScale(40, 120),                                          // BulletPenetration
	{12, 16, 20, 24, 28, 33, 37, 45, 55},                          // BulletDamage
	{15, 14, 13, 12, 11, 10, 9, 8, 7},                                // Reload
	linearScale(6, 14),                                            // MovementSpeed
}
