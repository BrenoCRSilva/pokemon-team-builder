package api

type Pokemon struct {
	Name      string           `json:"name"`
	Abilities []PokemonAbility `json:"abilities"`
	Sprites   PokemonSprite    `json:"sprites"`
	Stats     []PokemonStat    `json:"stats"`
	Types     []PokemonType    `json:"types"`
}

type PokemonAbility struct {
	Ability AbilityDetail `json:"ability"`
}

type AbilityDetail struct {
	Name string `json:"name"`
}

type PokemonSprite struct {
	FrontDefault string `json:"front_default"`
}

type PokemonStat struct {
	Stat     StatDetail `json:"stat"`
	BaseStat int        `json:"base_stat"`
}

type StatDetail struct {
	Name string `json:"name"`
}

type PokemonType struct {
	Details TypeDetail `json:"type"`
}

type TypeDetail struct {
	Name            string        `json:"name"`
	DamageRelations TypeRelations `json:"damage_relations"`
}

type TypeRelations struct {
	NoDamageTo       []TypeDetail `json:"no_damage_to"`
	HalfDamageTo     []TypeDetail `json:"half_damage_to"`
	DoubleDamageTo   []TypeDetail `json:"double_damage_to"`
	NoDamageFrom     []TypeDetail `json:"no_damage_from"`
	HalfDamageFrom   []TypeDetail `json:"half_damage_from"`
	DoubleDamageFrom []TypeDetail `json:"double_damage_from"`
}
