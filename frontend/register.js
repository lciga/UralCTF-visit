let participantCount = 0;
const maxParticipants = 7;
let cityChoices;
let uniChoices;

/**
 * Загрузка списка городов по поисковому запросу
 * @param {string} query - поисковый запрос для фильтрации городов
 */
async function loadCities(query = '') {
    const res = await fetch(`http://localhost:8080/api/cities/search?query=${encodeURIComponent(query)}`);
    if (!res.ok) return;
    const cities = await res.json();
    // Populate city select with search-enabled dropdown
    const cityOptions = cities.map(c => ({ value: c, label: c }));
    cityChoices.setChoices(cityOptions, 'value', 'label', true);
    // Load universities for first city by default
    if (cities.length > 0) {
        loadUniversities(cities[0]);
    }
}

async function loadUniversities(city) {
    const res = await fetch(`http://localhost:8080/api/search/university?city=${encodeURIComponent(city)}`);
    if (!res.ok) return;
    const unis = await res.json();
    // Populate university select based on selected city
    const uniOptions = unis.map(u => ({ value: u.id, label: u.name }));
    uniChoices.setChoices(uniOptions, 'value', 'label', true);
}

function addParticipant() {
    if (participantCount >= maxParticipants) return;
    participantCount++;
    const section = document.getElementById('participantsSection');
    const fieldset = document.createElement('fieldset');
    fieldset.innerHTML = `
        <legend>Участник №${participantCount}${participantCount === 1 ? ' (Капитан)' : ''}</legend>
        <label>Фамилия: <input name="participants[${participantCount-1}].last_name" required></label><br>
        <label>Имя: <input name="participants[${participantCount-1}].first_name" required></label><br>
        <label>Отчество: <input name="participants[${participantCount-1}].middle_name"></label><br>
        <label>Курс: <input type="number" name="participants[${participantCount-1}].course" min="1" max="6" required></label><br>
        <label>Telegram: <input name="participants[${participantCount-1}].telegram" required></label><br>
        <label>Email: <input type="email" name="participants[${participantCount-1}].email" required></label><br>
        <label>Размер футболки: <select name="participants[${participantCount-1}].shirt_size" required>
            <option value="S">S</option><option value="M">M</option><option value="L">L</option>
            <option value="XL">XL</option><option value="XXL">XXL</option>
        </select></label><br>
    `;
    section.appendChild(fieldset);
}

document.getElementById('addParticipant').addEventListener('click', () => {
    if (participantCount < 2) {
        addParticipant();
    } else {
        addParticipant();
    }
});

// Listen for city selection changes
document.addEventListener('change', function(e) {
    if (e.target && e.target.id === 'citySelect') {
        const selectedCity = cityChoices.getValue(true);
        loadUniversities(selectedCity);
    }
});

document.getElementById('teamForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    const form = e.target;
    const data = {
        name: form.name.value,
        city: form.city.value,
        university_id: parseInt(form.university_id.value),
        participants: []
    };
    for (let i = 0; i < participantCount; i++) {
        data.participants.push({
            first_name: form[`participants[${i}].first_name`].value,
            last_name: form[`participants[${i}].last_name`].value,
            middle_name: form[`participants[${i}].middle_name`].value,
            course: parseInt(form[`participants[${i}].course`].value),
            telegram: form[`participants[${i}].telegram`].value,
            email: form[`participants[${i}].email`].value,
            shirt_size: form[`participants[${i}].shirt_size`].value,
            is_captain: i === 0
        });
    }
    const res = await fetch('http://localhost:8080/api/teams', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(data)
    });
    if (res.ok) {
        alert('Команда зарегистрирована');
        window.location.href = 'index.html';
    } else {
        alert('Ошибка при регистрации');
    }
});

window.addEventListener('DOMContentLoaded', () => {
    // Initialize searchable selects
    const cityEl = document.getElementById('citySelect');
    const uniEl = document.getElementById('universitySelect');
    cityChoices = new Choices(cityEl, {
        searchEnabled: true,
        shouldSort: false,
        placeholderValue: 'Выберите город',
        callbackOnSearch: (searchTerm) => {
            loadCities(searchTerm);
        }
    });
    uniChoices = new Choices(uniEl, { searchEnabled: true, shouldSort: true, placeholderValue: 'Выберите университет' });
    // Load initial data and setup dynamic search
    loadCities();
    addParticipant();
    addParticipant();
});
