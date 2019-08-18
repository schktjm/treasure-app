import {storiesOf} from '@storybook/vue'

import LoginBtn from '../components/LoginBtn'

storiesOf('LoginBtn', module)
    .add('simple', () => ({
        components: {LoginBtn},
        template: `<LoginBtn>KEEP IT SIMPLE</LoginBtn>`
    }))
