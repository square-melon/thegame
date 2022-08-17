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
	{0, 0, 1, 2, 3, 4, 5, 6, 7},    // HealthRegen
	linearScale(1000, 5520),        // MaxHealth
	linearScale(15, 55),            // BodyDamage
	linearScale(8, 24),             // BulletSpeed
	linearScale(40, 136),           // BulletPenetration
	linearScale(18, 58),            // BulletDamage
	{15, 13, 11, 9, 7, 6, 5, 4, 3}, // Reload
	linearScale(6, 14),             // MovementSpeed
}
