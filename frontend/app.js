async function loadTeams() {
    const res = await fetch('http://localhost:8080/api/teams');
    if (!res.ok) return;
    const teams = await res.json();
    const container = document.getElementById('teams');
    container.innerHTML = teams.map(t => `
        <div class="team">
            <h2>${t.name}</h2>
            <p>Город: ${t.city_name}</p>
            <p>Университет: ${t.university_name}</p>
        </div>
    `).join('');
}

window.addEventListener('DOMContentLoaded', loadTeams);
