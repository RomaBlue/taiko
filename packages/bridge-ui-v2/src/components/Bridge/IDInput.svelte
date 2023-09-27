<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import { t } from 'svelte-i18n';

  import { FlatAlert } from '$components/Alert';
  import { Icon } from '$components/Icon';

  export let numbersArray: number[] = [];

  enum State {
    Valid,
    Invalid,
    Neutral,
  }

  const dispatch = createEventDispatcher();
  let inputId = 'numberInput';
  let state = State.Neutral;
  let value: string = '';

  $: {
    value = value.replace(/\s+/g, '');
    if (value === '' || value.endsWith(',')) {
      state = State.Neutral;
      numbersArray = [];
    } else {
      const inputArray = value.split(',');
      const isValid = inputArray.every((item) => /^[0-9]+$/.test(item));
      state = isValid ? State.Valid : State.Invalid;
      numbersArray = isValid ? inputArray.map((num) => parseInt(num)).filter(Boolean) : [];
    }
  }

  function validateInput(e: Event) {
    const target = e.target as HTMLInputElement;
    value = target.value.replace(/\s+/g, '');
    dispatch('input', { value, numbersArray });
  }

  function clear() {
    value = '';
    numbersArray = [];
    state = State.Neutral;
    dispatch('input', { value, numbersArray });
  }
</script>

<div class="f-col space-y-2">
  <div class="f-between-center text-secondary-content">
    <label class="body-regular" for={inputId}>{$t('inputs.nft.token_id.label')}</label>
  </div>
  <div class="relative f-items-center">
    <input
      id={inputId}
      type="text"
      placeholder={$t('inputs.nft.token_id.placeholder')}
      bind:value
      on:input={validateInput}
      class="w-full input-box withValdiation py-6 pr-16 px-[26px] title-subsection-bold placeholder:text-tertiary-content
      {state === State.Valid ? 'success' : state === State.Invalid ? 'error' : ''}" />
    <button class="absolute right-6 uppercase body-bold text-secondary-content" on:click={clear}>
      <Icon type="x-close-circle" fillClass="fill-primary-icon" size={24} />
    </button>
  </div>
</div>

<div class="mt-5 min-h-[20px]">
  {#if state === State.Invalid && value !== ''}
    <FlatAlert type="error" forceColumnFlow message={$t('inputs.address_input.errors.invalid')} />
  {:else if state === State.Valid}
    <FlatAlert type="success" forceColumnFlow message={$t('inputs.address_input.success')} />
  {/if}
</div>
