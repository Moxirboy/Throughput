package db

// Client represents a client entity
type Client struct {
	Name string
	Date string
}

func GetClientNames() []string {
	rows, err := DB.Query("SELECT name FROM kirim.client;")
	if err != nil {
		return nil
	}
	defer rows.Close()
	var clientNames []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil
		}
		clientNames = append(clientNames, name)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return clientNames
}
func GenerateDropdownHTMLClient() string {
	clientNames := GetClientNames()
	dropdownHTML := "<select name='client'>"
	for _, name := range clientNames {
		dropdownHTML += "<option value='" + name + "'>" + name + "</option>"
	}

	dropdownHTML += "</select>"
	return dropdownHTML
}
