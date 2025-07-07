document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('order-form');
    form.addEventListener('submit', findOrder);
});

async function findOrder(event) {
    event.preventDefault();

    const input = document.getElementById('order-id-input');
    const errorSpan = document.getElementById('form-error');
    const detailsDiv = document.getElementById('details');
    errorSpan.textContent = '';
    detailsDiv.innerHTML = '';

    let uid = input.value.trim();
    if (!uid) {
        errorSpan.textContent = 'Введите Order UID';
        return;
    }

    try {
        let res = await fetch('/order/' + encodeURIComponent(uid));
        if (!res.ok) {
            errorSpan.textContent = 'Заказ не найден';
            return;
        }
        let data = await res.json();
        renderOrderDetails(data);
    } catch (e) {
        errorSpan.textContent = 'Ошибка запроса';
    }
}

function renderOrderDetails(data) {
    let date = new Date(data.date_created).toLocaleString('ru-RU', {
        day: '2-digit', month: '2-digit', year: 'numeric',
        hour: '2-digit', minute: '2-digit'
    });

    let payment = data.payment;
    let delivery = data.delivery;
    let items = data.items;

    // Делаем уникальные id для input/label чтобы не было конфликтов
    let uid = data.order_uid.replace(/[^a-zA-Z0-9]/g, '').slice(-8);
    let html = `
    <h2>Заказ <span style="color:#3af">${data.order_uid}</span></h2>
    <div class="accordion">
      <div class="accordion-section">
        <input type="checkbox" id="payment-accordion-${uid}" checked>
        <label for="payment-accordion-${uid}">💳 Оплата</label>
        <div class="content">
          <table>
            <tr><td class="field">Сумма</td><td class="value">${payment.amount} ${getCurrencyIcon(payment.currency)}</td></tr>
            <tr><td class="field">Провайдер</td><td class="value">${payment.provider}</td></tr>
            <tr><td class="field">Банк</td><td class="value">${payment.bank}</td></tr>
            <tr><td class="field">Доставка</td><td class="value">${payment.delivery_cost}</td></tr>
            <tr><td class="field">Товары</td><td class="value">${payment.goods_total}</td></tr>
            <tr><td class="field">Custom Fee</td><td class="value">${payment.custom_fee}</td></tr>
          </table>
        </div>
      </div>
      <div class="accordion-section">
        <input type="checkbox" id="delivery-accordion-${uid}">
        <label for="delivery-accordion-${uid}">🚚 Доставка</label>
        <div class="content">
          <table>
            <tr><td class="field">Имя</td><td class="value">${delivery.name}</td></tr>
            <tr><td class="field">Город</td><td class="value">${delivery.city}</td></tr>
            <tr><td class="field">Адрес</td><td class="value">${delivery.address}</td></tr>
            <tr><td class="field">Регион</td><td class="value">${delivery.region}</td></tr>
            <tr><td class="field">Индекс</td><td class="value">${delivery.zip}</td></tr>
            <tr><td class="field">Телефон</td><td class="value">${delivery.phone}</td></tr>
            <tr><td class="field">Email</td><td class="value">${delivery.email}</td></tr>
          </table>
        </div>
      </div>
      <div class="accordion-section">
        <input type="checkbox" id="items-accordion-${uid}">
        <label for="items-accordion-${uid}">📦 Товары</label>
        <div class="content">
          <ul>
            ${items.map(x => `
              <li>
                <b>${x.name}</b> <span style="color:#88e">[${x.price}×${x.size}]</span>
                <br><span style="color:#bbb">Бренд: ${x.brand}, Артикул: ${x.nm_id}</span>
              </li>
            `).join('')}
          </ul>
        </div>
      </div>
      <div class="accordion-section">
        <input type="checkbox" id="info-accordion-${uid}">
        <label for="info-accordion-${uid}">ℹ️ Прочее</label>
        <div class="content">
          <table>
            <tr><td class="field">Трек</td><td class="value">${data.track_number}</td></tr>
            <tr><td class="field">Дата создания</td><td class="value">${date}</td></tr>
            <tr><td class="field">Клиент</td><td class="value">${data.customer_id}</td></tr>
            <tr><td class="field">Сервис</td><td class="value">${data.delivery_service}</td></tr>
            <tr><td class="field">ShardKey</td><td class="value">${data.shardkey}</td></tr>
            <tr><td class="field">SmID</td><td class="value">${data.sm_id}</td></tr>
            <tr><td class="field">OofShard</td><td class="value">${data.oof_shard}</td></tr>
            <tr><td class="field">Локаль</td><td class="value">${data.locale}</td></tr>
          </table>
        </div>
      </div>
    </div>
    `;
    document.getElementById('details').innerHTML = html;
}

function getCurrencyIcon(currency) {
    if (currency === "USD") return "💵";
    if (currency === "RUB") return "₽";
    if (currency === "EUR") return "€";
    return currency;
}
