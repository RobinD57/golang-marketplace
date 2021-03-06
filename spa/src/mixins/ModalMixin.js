export default {
  methods: {
    showModal() {
      this.$store.commit('setModalOpen', true);
      if (this.modalOpen) {
        document.querySelector('.modal-overlay').style.display = "block";
        this.$refs.modal.style.display = "";
        if (this.$options.name == 'PurchaseButton') {
          injectCardInModal(this.$refs.cWrap, this.listing._id)
        }
      }
    },
    removeOverlay() {
      this.$store.commit('setModalOpen', false);
      this.working = false;
      document.querySelector('.modal-overlay').style.display = 'none';
      document.querySelectorAll('.modal').forEach((modal) => {
        modal.style.display = 'none';
      });
        this.next = false;
    },
  }
}

const injectCardInModal = (target, id) => {
  const cardHTML = document.querySelector(`[data-id="${id}"]`).innerHTML;
  target.innerHTML = cardHTML;
}
