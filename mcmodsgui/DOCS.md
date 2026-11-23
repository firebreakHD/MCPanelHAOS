# Einrichtung

1. Die App installieren. In der App-Info den Schutzmodus für **MCPanelHAOS** ausschalten. Die App erhält dadurch Zugriff auf Docker wie Portainer.
2. Falls dein HAOS unter einem anderen Namen erreichbar ist, `panel_hostname` auf seinen DNS-Namen oder seine IPv4-Adresse setzen. Die Vorgabe lautet `homeassistant.local`. Den HTTPS-Port 8443 in der App-Netzwerkkonfiguration auf 8443 lassen oder bei Änderung auch `panel_port` anpassen.
3. Starten, die Seitenleiste aktivieren und **MCPanelHAOS** öffnen. Beim ersten Aufruf das Stammzertifikat aus der Home-Assistant-Seite beziehen und auf deinem Verwaltungsgerät als vertrauenswürdig hinterlegen. Anschließend **Minecraft-Verwaltung öffnen** und das Administrator-Konto im eigenen LAN einrichten.
4. Zwei Java-Profile mit den IDs `modpack-a` und `modpack-b` anlegen. Servernamen frei wählen. In der Oberfläche CurseForge konfigurieren und die Packs auswählen.
5. Das gewünschte Profil starten. Beim Start des anderen Profils fragt das Panel, ob der aktive Server gespeichert und sauber gestoppt werden soll.

Spieler verbinden sich mit der Adresse deines Home-Assistant-Geräts und Port 25565/TCP. Dieser Port wird direkt vom jeweils laufenden Minecraft-Container veröffentlicht. Die Verwaltung auf 8443 bleibt im LAN oder VPN.

## Updates und Backups

Vor einem HAOS-App-Backup oder Update zuerst über das Panel Minecraft stoppen und den Abschluss abwarten. Die App versucht bei einem regulären Stopp zusätzlich, einen aktiven Server geordnet zu beenden. Geht ein Stopp schief, im Panel den Minecraft-Zustand prüfen. Welten, Panel-Datenbank, TLS-CA und Sitzungsschlüssel liegen im privaten, persistenten Mount dieser App. HA-Backups der App erfassen diese Daten und enthalten sensible Zugangsdaten. Backups entsprechend geschützt aufbewahren.
