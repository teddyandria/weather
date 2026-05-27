# weather

| Donnée               | Comment c'est représenté en JSON ? | Comment c'est représenté en XML ?                                          |
|:---------------------|:-----------------------------------|:---------------------------------------------------------------------------|
| Pays                 | Champ `"country"` avec le nom complet (ex : `"France"`) | Attribut `country` avec le code ISO 2 lettres (ex : `country="FR"`)        |
| Coordonnées          | Objet imbriqué `"location"` avec les clés `"latitude"` et `"longitude"` | Attributs `lat` et `lon` sur l'élément `<coordinates>`                     |
| Altitude             | Champ `"altitude_m"` au niveau de la station (ex : `"altitude_m": 47`) | Attribut `altitude` sur l'élément `<coordinates>` (ex : `altitude="47"`)   |
| Modèle de capteur    | Objet `"device"` avec les clés `"type"`, `"manufacturer"` et `"installed_on"` | Élément `<hardware>` avec les attributs `vendor`, `model` et `since`       |
| Température          | Champ nommé `"temperature_celsius"` (ex : `"temperature_celsius": 3.3`) | Élément générique `<measure type="temperature" unit="C">3.3</measure>`     |
| Conditions ciel      | Champ `"conditions"` dans l'observation (ex : `"conditions": "clear"`) | Attribut `sky` sur l'élément `<observation>` (ex : `sky="clear"`)          |
| Vent                 | Objet imbriqué `"wind"` avec les clés `"speed_kmh"` et `"direction_deg"` | Élément auto-fermant `<wind>` avec les attributs `speed` et `direction`    |
| Notes (optionnelles) | Champ `"notes"` toujours présent, vaut `null` si absent | Élément `<note>` présent uniquement quand il y a une valeur (absent sinon) |
