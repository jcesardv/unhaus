package model

func GetAllUnhaus() ([]Unhaus, error) {
	var unhaus []Unhaus
	tx := db.Find(&unhaus)
	if tx.Error != nil {
		return []Unhaus{}, tx.Error
	}
	return unhaus, nil
}

func GetUnhaus(id uint64) (Unhaus, error) {
	var unhaus Unhaus
	tx := db.Where("id = ?",  id).First(&unhaus)
	if tx.Error != nil{
		return Unhaus{}, tx.Error
	}
	return unhaus, nil
}

func CreateUnhaus(unhaus Unhaus) error {
	tx := db.Create(&unhaus)
	return tx.Error
}

func UpdateUnhaus(unhaus Unhaus) error {
	tx := db.Save(&unhaus)
	return tx.Error
}

func DeleteUnhaus(id uint64) error {
	tx := db.Unscoped().Delete(&Unhaus{}, id)
	return tx.Error
}

func FindByUnhausUrl(url string) (Unhaus, error) {
	var unhaus Unhaus
	tx := db.Where("unhaus = ?", url).First(&unhaus)
	return unhaus, tx.Error
}