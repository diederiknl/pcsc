# Doel van dit document

Ok, dit is wellicht iets overdreven, maar ik denk dat het goed is voor mijzelf om frustraties en denkwijzes duidelijk te kunnen maken.
Waarom heb ik bepaalde keuzes overwogen?

Voor de lezer: ik ben multigeinteresseerd. En wil meerdere dingen tegelijkertijd leren. Dat brengt mij nog wel eens in de problemen ;)

## 30 oktober 2024

Ik heb de kaartlezer heel even laten zitten. Ik zat mij te bedenken dat dit probleem eigenlijk om een vreemde manier aanvlieg. Een alternatieve aanpak zou kunnen zijn dat docenten een 'inlogtoken' aanvragen op bijvoorbeeld inchecken.diederik.nl, en dat ze daarbij authenticeren per email.
Vervolgens zou via JavaScript een refresh plaats kunnen vinden naar een QR-Code waar studenten weer in kunnen checken. Geen kaartlezer nodig dus :)

In de directory API heb ik door Copilot een API laten genereren. Kenmerkend daarbij is dat ik ook een reader-ID mee stuur. Rationale hierachter is dat we mogelijk meerdere kaartlezers hebben.

Vervolgens ben ik tegen de structuur van GO aan gelopen. API is wat mij betreft nu geen package, maar een hoofdstructuur. Hoe Git hier mee om gaat is voor mij nog een vraagteken.

-[ ] Hoe werkt Git met packages?