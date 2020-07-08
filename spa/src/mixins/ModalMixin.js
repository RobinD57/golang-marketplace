export default {
  methods: {
    showModal() {
      this.modalOpen = true;
      if (this.modalOpen) {
        document.querySelector('.modal-overlay').style.display = "block";
        this.$refs.modal.style.display = "";
        if (this.$options.name == 'PurchaseButton') {
          this.injectCardInModal();
        }
      }
    },
    removeOverlay(attrs) {
      this.modalOpen = false;
      document.querySelector('.modal-overlay').style.display = 'none';
      document.querySelectorAll('.modal').forEach((modal) => {
        modal.style.display = 'none';
      });
      attrs.forEach((attr) => {
        this[attr] = false;
        attr;
      });

    },
    injectCardInModal() {
      const cardHTML = document.querySelector(`[data-id="${this.listing._id}"]`).innerHTML;
      this.$refs.cWrap.innerHTML = cardHTML;
    }
  }
}
