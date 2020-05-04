import React from "react"
import PropTypes from "prop-types"
import { colors } from "../constants/colors"
import { ReactComponent as ArrowRight } from "../static/svg/arrow-right.svg"
import { ReactComponent as Operations } from "../static/svg/operations.svg"
import { ReactComponent as Rewards } from "../static/svg/rewards.svg"
import { ReactComponent as Glossary } from "../static/svg/glossary.svg"
import { ReactComponent as Authorizer } from "../static/svg/authorizer.svg"
import { ReactComponent as KeepToken } from "../static/svg/tokens.svg"
import { ReactComponent as GrantContextIcon } from "../static/svg/grant-context-icon.svg"
import { ReactComponent as MoneyWalletOpen } from "../static/svg/money-wallet-open.svg"
import { ReactComponent as KeepGreenOutline } from "../static/svg/keep-green-form-icon.svg"
import { ReactComponent as AuthorizerFormIcon } from "../static/svg/authorizer-form-icon.svg"
import { ReactComponent as OperatorFormIcon } from "../static/svg/operator-form-icon.svg"
import { ReactComponent as BeneficiaryFormIcon } from "../static/svg/beneficiary-form-icon.svg"
import { ReactComponent as DashedLine } from "../static/svg/dashed-line.svg"

const Keep = ({ color, height, width }) => (
  <svg
    width={width}
    height={height}
    viewBox="0 0 718 126"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      d="M63 0C28.2176 0 0 28.2176 0 63C0 97.7824 28.2176 126 63 126C97.7824 126 126 97.7824 126 63C125.947 28.2176 97.7294 0 63 0ZM89.4176 48.9176H85.4471L73.4294 62.9471L85.4471 76.9765H89.4176V88.2529H63.9V77.0824H67.8706L59.9824 67.8706H56.7529V77.0824H61.2529V88.2529H36.5294V76.9765H41.8765V63V48.9176H36.5294V37.6412H42.6176V41.8235H45.9V37.6412H51.8824V41.8235H55.1647V37.6412H61.2V48.7588H56.7V57.9706H59.9294L67.8176 48.7588H63.8471V37.6412H89.3647V48.9176H89.4176Z"
      fill="#0A0806"
    />
    <path
      d="M179.59 79.66H173.22C171.12 79.66 170.35 79.38 169.93 78.96C169.37 78.4 169.16 77.56 169.16 75.6V51.03H179.66V46.55H169.16V37.17H164.12V46.55H158.31V51.03H164.12V76.65C164.12 79.66 164.47 81.27 165.73 82.53C167.06 83.86 168.67 84.28 171.12 84.28C173.15 84.28 174.76 84 176.09 83.79L179.59 82.95V79.66Z"
      fill="black"
    />
    <path
      d="M197.937 84H207.317C211.167 84 213.337 83.51 214.947 81.9C217.327 79.52 218.097 76.51 218.097 65.24C218.097 54.04 217.327 51.03 214.947 48.65C213.337 47.04 211.167 46.55 207.317 46.55H197.937C194.087 46.55 191.917 47.04 190.307 48.65C187.927 51.03 187.157 54.04 187.157 65.24C187.157 76.51 187.927 79.52 190.307 81.9C191.917 83.51 194.087 84 197.937 84ZM192.267 65.24C192.267 55.16 192.827 53.34 193.947 52.22C194.787 51.38 195.837 51.1 198.497 51.1H206.757C209.417 51.1 210.467 51.38 211.307 52.22C212.427 53.34 212.987 55.16 212.987 65.24C212.987 75.39 212.427 77.21 211.307 78.33C210.467 79.17 209.417 79.45 206.757 79.45H198.497C195.837 79.45 194.787 79.17 193.947 78.33C192.827 77.21 192.267 75.39 192.267 65.24Z"
      fill="black"
    />
    <path
      d="M257.446 50.96H262.766V46.55H246.036V50.89H251.846L241.626 62.93H236.586V33.6H225.386V38.08H231.546V79.59H225.386V84H242.466V79.59H236.586V67.27H241.416L253.736 84H263.186V79.59H256.956L245.546 64.68L257.446 50.96Z"
      fill="black"
    />
    <path
      d="M298.11 67.13C298.74 67.13 298.95 66.85 298.95 66.29V65.03C298.95 54.25 298.18 51.03 295.73 48.58C294.26 47.11 291.95 46.55 288.17 46.55H279.91C276.06 46.55 273.89 47.04 272.28 48.65C269.9 51.03 269.13 54.04 269.13 65.24C269.13 76.51 269.9 79.52 272.28 81.9C273.89 83.51 276.06 84 279.91 84H289.57C293.21 84 294.96 83.44 296.36 82.04C298.18 80.22 298.81 78.26 298.95 72.52H293.98C293.84 76.37 293.49 77.56 292.72 78.4C291.88 79.17 290.97 79.52 288.59 79.52H280.4C277.74 79.52 276.76 79.31 275.92 78.4C274.87 77.42 274.31 75.67 274.24 67.13H298.11ZM275.85 52.15C276.76 51.31 277.74 51.03 280.4 51.03H287.68C290.34 51.03 291.39 51.31 292.16 52.15C293.21 53.2 293.84 55.23 293.84 62.93H274.24C274.31 55.02 274.87 53.2 275.85 52.15Z"
      fill="black"
    />
    <path
      d="M317.89 46.55H306.69V51.03H312.85V79.59H306.69V84H323.77V79.59H317.89V54.39C321.81 51.87 323.91 51.03 327.06 51.03H330.91C333.64 51.03 334.41 51.31 335.18 52.08C335.95 52.92 336.3 54.04 336.3 56.21V79.59H330.35V84H347.36V79.59H341.27V55.93C341.27 51.94 340.78 49.91 339.31 48.44C337.91 47.11 336.37 46.48 332.87 46.48H328.53C324.54 46.48 322.79 47.53 317.89 51.52V46.55Z"
      fill="black"
    />
    <path
      d="M378.691 35H352.091V39.69H358.951V79.31H352.091V84H378.691C383.661 84 386.251 83.09 388.421 80.99C391.221 78.12 392.971 72.66 392.971 59.5C392.971 46.34 391.221 40.88 388.421 38.01C386.251 35.91 383.661 35 378.691 35ZM364.201 39.69H378.341C381.491 39.69 382.821 40.04 384.221 41.44C386.111 43.26 387.371 47.25 387.371 59.57C387.371 71.75 386.111 75.74 384.221 77.56C382.891 78.96 381.491 79.31 378.341 79.31H364.201V39.69Z"
      fill="black"
    />
    <path
      d="M426.92 57.33V63.42C423.21 62.79 419.29 62.58 415.93 62.58C409.42 62.58 406.83 63.49 405.22 65.1C403.75 66.57 402.91 69.23 402.91 73.71C402.91 78.47 403.82 80.64 405.15 81.97C406.69 83.51 408.79 84.07 412.15 84.07H416.56C420.27 84.07 422.09 83.16 426.92 79.24V84H437.98V79.59H431.96V57.05C431.96 52.15 431.61 50.12 429.93 48.51C428.67 47.18 426.99 46.55 422.93 46.55H413.83C410.33 46.55 408.51 47.04 406.97 48.58C405.43 50.19 404.73 52.22 404.59 57.75H409.63C409.7 54.39 410.05 52.92 410.96 52.01C411.73 51.24 412.92 50.96 415.3 50.96H422.02C424.54 50.96 425.31 51.24 425.94 51.87C426.64 52.57 426.92 53.62 426.92 57.33ZM417.33 79.66H413.27C411.1 79.66 409.98 79.38 409.14 78.61C408.37 77.84 408.02 76.3 408.02 73.01C408.02 69.93 408.37 68.81 409.28 67.83C410.19 66.92 411.38 66.5 415.72 66.5H426.92V76.3C422.79 78.89 420.69 79.66 417.33 79.66Z"
      fill="black"
    />
    <path
      d="M444.699 56.84C444.699 61.67 445.189 63.49 446.589 64.89C447.639 65.94 449.179 66.43 452.609 66.78L463.389 67.9C465.769 68.18 466.469 68.46 466.959 68.95C467.519 69.51 467.869 70.49 467.869 73.71C467.869 77.28 467.449 78.19 466.819 78.82C466.259 79.38 465.419 79.66 463.739 79.66H453.239C451.279 79.66 450.509 79.38 450.019 78.89C449.319 78.19 449.109 77.21 448.969 72.66H444.209C444.279 79.24 444.769 80.78 446.379 82.39C447.429 83.44 449.459 84 452.329 84H464.929C467.659 84 469.269 83.51 470.459 82.32C471.999 80.71 472.769 78.54 472.769 73.43C472.769 68.74 472.349 66.57 470.949 65.17C469.969 64.19 468.359 63.7 464.999 63.35L454.079 62.16C451.629 61.95 450.999 61.67 450.369 61.11C449.809 60.48 449.599 59.57 449.599 56.56C449.599 53.34 450.019 52.36 450.719 51.66C451.279 51.1 452.119 50.89 453.659 50.89H463.249C465.069 50.89 465.909 51.24 466.399 51.73C467.029 52.36 467.379 53.2 467.519 57.61H472.279C472.209 51.38 471.579 49.63 470.109 48.16C469.059 47.11 467.379 46.55 464.369 46.55H452.539C449.739 46.55 448.199 47.04 447.009 48.23C445.469 49.77 444.699 52.15 444.699 56.84Z"
      fill="black"
    />
    <path
      d="M491.05 33.6H479.85V38.08H486.01V79.59H479.85V84H496.93V79.59H491.05V54.39C494.97 51.87 497.14 51.03 500.22 51.03H504.07C506.8 51.03 507.57 51.31 508.34 52.08C509.18 52.92 509.46 54.04 509.46 56.21V79.59H503.58V84H520.52V79.59H514.5V55.93C514.5 51.94 513.94 49.91 512.54 48.44C511.14 47.11 509.53 46.48 506.1 46.48H501.69C497.77 46.48 496.02 47.53 491.05 51.52V33.6Z"
      fill="black"
    />
    <path
      d="M535.205 33.6H524.005V38.08H530.165V79.59H524.005V84H535.205V79.03C540.175 83.09 541.995 84.07 545.845 84.07H549.765C553.965 84.07 556.135 83.37 557.745 81.76C559.915 79.52 560.685 76.58 560.685 65.24C560.685 53.97 559.915 51.03 557.745 48.79C556.135 47.18 553.965 46.48 549.765 46.48H545.845C541.995 46.48 540.175 47.46 535.205 51.52V33.6ZM544.375 51.03H548.715C551.725 51.03 552.915 51.45 553.755 52.36C555.155 53.69 555.505 56 555.505 65.24C555.505 74.55 555.155 76.79 553.755 78.19C552.915 79.03 551.725 79.52 548.715 79.52H544.585C541.295 79.52 539.055 78.54 535.205 76.09V54.39C538.915 52.01 541.155 51.03 544.375 51.03Z"
      fill="black"
    />
    <path
      d="M581.137 84H590.517C594.367 84 596.537 83.51 598.147 81.9C600.527 79.52 601.297 76.51 601.297 65.24C601.297 54.04 600.527 51.03 598.147 48.65C596.537 47.04 594.367 46.55 590.517 46.55H581.137C577.287 46.55 575.117 47.04 573.507 48.65C571.127 51.03 570.357 54.04 570.357 65.24C570.357 76.51 571.127 79.52 573.507 81.9C575.117 83.51 577.287 84 581.137 84ZM575.467 65.24C575.467 55.16 576.027 53.34 577.147 52.22C577.987 51.38 579.037 51.1 581.697 51.1H589.957C592.617 51.1 593.667 51.38 594.507 52.22C595.627 53.34 596.187 55.16 596.187 65.24C596.187 75.39 595.627 77.21 594.507 78.33C593.667 79.17 592.617 79.45 589.957 79.45H581.697C579.037 79.45 577.987 79.17 577.147 78.33C576.027 77.21 575.467 75.39 575.467 65.24Z"
      fill="black"
    />
    <path
      d="M634.416 57.33V63.42C630.706 62.79 626.786 62.58 623.426 62.58C616.916 62.58 614.326 63.49 612.716 65.1C611.246 66.57 610.406 69.23 610.406 73.71C610.406 78.47 611.316 80.64 612.646 81.97C614.186 83.51 616.286 84.07 619.646 84.07H624.056C627.766 84.07 629.586 83.16 634.416 79.24V84H645.476V79.59H639.456V57.05C639.456 52.15 639.106 50.12 637.426 48.51C636.166 47.18 634.486 46.55 630.426 46.55H621.326C617.826 46.55 616.006 47.04 614.466 48.58C612.926 50.19 612.226 52.22 612.086 57.75H617.126C617.196 54.39 617.546 52.92 618.456 52.01C619.226 51.24 620.416 50.96 622.796 50.96H629.516C632.036 50.96 632.806 51.24 633.436 51.87C634.136 52.57 634.416 53.62 634.416 57.33ZM624.826 79.66H620.766C618.596 79.66 617.476 79.38 616.636 78.61C615.866 77.84 615.516 76.3 615.516 73.01C615.516 69.93 615.866 68.81 616.776 67.83C617.686 66.92 618.876 66.5 623.216 66.5H634.416V76.3C630.286 78.89 628.186 79.66 624.826 79.66Z"
      fill="black"
    />
    <path
      d="M661.505 79.59V55.02C665.495 52.29 667.455 51.8 670.325 51.8H674.945V46.48H670.885C667.735 46.48 666.055 47.88 661.505 52.64V46.55H650.305V51.03H656.465V79.59H650.305V84H667.735V79.59H661.505Z"
      fill="black"
    />
    <path
      d="M706.593 84H717.793V79.59H711.633V33.6H700.153V38.01H706.593V51.52C701.693 47.53 699.803 46.48 695.953 46.48H692.033C687.833 46.48 685.733 47.18 684.123 48.79C681.883 51.03 681.183 53.97 681.183 65.24C681.183 76.58 681.883 79.52 684.123 81.76C685.733 83.37 687.833 84.07 692.033 84.07H695.953C699.803 84.07 701.623 83.09 706.593 79.03V84ZM697.423 79.52H693.083C690.143 79.52 688.953 79.03 688.043 78.19C686.713 76.79 686.293 74.55 686.293 65.24C686.293 56 686.713 53.69 688.043 52.36C688.953 51.45 690.143 51.03 693.083 51.03H697.283C700.503 51.03 702.743 52.01 706.593 54.53V76.16C702.813 78.61 700.573 79.52 697.423 79.52Z"
      fill="black"
    />
  </svg>
)

Keep.propTypes = {
  color: PropTypes.string,
  height: PropTypes.string,
  width: PropTypes.string,
}

Keep.defaultProps = {
  color: "#293330",
  height: "238.16",
  width: "917.41333",
}

const Badge = ({ height, width }) => (
  <svg
    width={width}
    height={height}
    viewBox="0 0 55 55"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      d="M11.6155 29.7993L2.65625 40.7482L10.9375 42.4045L14.25 52.342L22.9597 38.5222"
      stroke="black"
      strokeWidth="4"
      strokeLinecap="round"
      strokeLinejoin="round"
    />
    <path
      d="M43.3846 29.7993L52.3438 40.7482L44.0626 42.4045L40.7501 52.342L32.0404 38.5222"
      stroke="black"
      strokeWidth="4"
      strokeLinecap="round"
      strokeLinejoin="round"
    />
    <path
      fillRule="evenodd"
      clipRule="evenodd"
      d="M27.5 39.0918C37.5619 39.0918 45.7187 30.935 45.7187 20.873C45.7187 10.8111 37.5619 2.6543 27.5 2.6543C17.4381 2.6543 9.28125 10.8111 9.28125 20.873C9.28125 30.935 17.4381 39.0918 27.5 39.0918Z"
      stroke="black"
      strokeWidth="4"
      strokeLinecap="round"
      strokeLinejoin="round"
    />
    <path
      fillRule="evenodd"
      clipRule="evenodd"
      d="M30.825 22.5293C30.825 24.3587 29.3419 25.8418 27.5125 25.8418C25.683 25.8418 24.2 24.3587 24.2 22.5293V19.2168C24.2 17.3874 25.683 15.9043 27.5125 15.9043C29.3419 15.9043 30.825 17.3874 30.825 19.2168V22.5293Z"
      stroke="black"
      strokeWidth="4"
      strokeLinecap="round"
      strokeLinejoin="round"
    />
  </svg>
)

const Cross = ({ height, width, ...restProps }) => (
  <svg
    height={height}
    width={width}
    {...restProps}
    viewBox="0 0 20 20"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      d="M1 19L19 1"
      stroke={restProps.color}
      strokeLinecap="round"
      strokeLinejoin="round"
    />
    <path
      d="M19 19L1 1"
      stroke={restProps.color}
      strokeLinecap="round"
      strokeLinejoin="round"
    />
  </svg>
)

Cross.defaultProps = {
  color: "#6D6D6D",
}

const Tooltip = ({ backgroundColor, color }) => (
  <svg
    width="15"
    height="16"
    viewBox="0 0 15 16"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      fillRule="evenodd"
      clipRule="evenodd"
      d="M13.117 12.3021H8.3426L4.32209 15.2888V12.3021H2.31183C1.75671 12.3021 1.3067 11.8564 1.3067 11.3066V1.35102C1.3067 0.801194 1.75671 0.355469 2.31183 0.355469H13.117C13.6721 0.355469 14.1221 0.801194 14.1221 1.35102V11.3066C14.1221 11.8564 13.6721 12.3021 13.117 12.3021Z"
      fill={backgroundColor}
    />
    <path
      fillRule="evenodd"
      clipRule="evenodd"
      d="M7.51076 4.33833C7.3878 3.80981 6.88395 3.45686 6.33987 3.5181C5.79578 3.57934 5.38474 4.03526 5.38464 4.57763C5.3846 4.774 5.22385 4.93316 5.02559 4.93312C4.82734 4.93308 4.66665 4.77386 4.66669 4.5775C4.66686 3.67351 5.35196 2.9136 6.2588 2.81153C7.16565 2.70946 8.00542 3.29774 8.21038 4.17864C8.41533 5.05954 7.92019 5.95251 7.05971 6.25384C6.91623 6.30408 6.82035 6.43847 6.82043 6.5892V6.71074C6.82043 6.9071 6.65971 7.06629 6.46145 7.06629C6.26319 7.06629 6.10248 6.9071 6.10248 6.71074L6.10248 6.58957C6.10248 6.58951 6.10248 6.58963 6.10248 6.58957C6.10232 6.13743 6.38999 5.73412 6.82039 5.5834C7.33666 5.40261 7.63373 4.86684 7.51076 4.33833Z"
      fill={color}
    />
    <path
      d="M6.10248 6.58957L6.10248 6.71074C6.10248 6.9071 6.26319 7.06629 6.46145 7.06629C6.65971 7.06629 6.82043 6.9071 6.82043 6.71074V6.5892C6.82035 6.43847 6.91623 6.30408 7.05971 6.25384C7.92019 5.95251 8.41533 5.05954 8.21038 4.17864C8.00542 3.29774 7.16565 2.70946 6.2588 2.81153C5.35196 2.9136 4.66686 3.67351 4.66669 4.5775C4.66665 4.77386 4.82734 4.93308 5.02559 4.93312C5.22385 4.93316 5.3846 4.774 5.38464 4.57763C5.38474 4.03526 5.79578 3.57934 6.33987 3.5181C6.88395 3.45686 7.3878 3.80981 7.51076 4.33833C7.63373 4.86684 7.33666 5.40261 6.82039 5.5834C6.38999 5.73412 6.10232 6.13743 6.10248 6.58957ZM6.10248 6.58957C6.10248 6.58963 6.10248 6.58951 6.10248 6.58957Z"
      strokeWidth="0.25"
      strokeLinecap="round"
      strokeLinejoin="round"
      stroke={color}
    />
    <path
      fillRule="evenodd"
      clipRule="evenodd"
      d="M5.92267 8.46667C5.92267 8.17213 6.16374 7.93335 6.46112 7.93335C6.75849 7.93335 6.99956 8.17213 6.99956 8.46667C6.99956 8.76121 6.75849 8.99999 6.46112 8.99999C6.16374 8.99999 5.92267 8.76121 5.92267 8.46667Z"
      fill={color}
      stroke={color}
      strokeWidth="0.1"
      strokeLinecap="round"
      strokeLinejoin="round"
    />
    <path
      fillRule="evenodd"
      clipRule="evenodd"
      d="M1.3052 0.705146C0.977543 0.705146 0.711926 0.968233 0.711926 1.29277V10.6947C0.711926 11.0192 0.977543 11.2823 1.3052 11.2823H3.20367C3.40026 11.2823 3.55963 11.4402 3.55963 11.6349V13.7503L6.78703 11.3528C6.84864 11.3071 6.92358 11.2823 7.0006 11.2823H11.5095C11.8371 11.2823 12.1027 11.0192 12.1027 10.6947V1.29277C12.1027 0.968233 11.8371 0.705146 11.5095 0.705146H1.3052ZM0 1.29277C0 0.578791 0.584357 0 1.3052 0H11.5095C12.2303 0 12.8147 0.578791 12.8147 1.29277V10.6947C12.8147 11.4087 12.2303 11.9875 11.5095 11.9875H7.11926L3.41724 14.7375C3.30938 14.8177 3.16507 14.8306 3.04448 14.7708C2.92388 14.7111 2.8477 14.589 2.8477 14.4555V11.9875H1.3052C0.584357 11.9875 0 11.4087 0 10.6947V1.29277Z"
      fill={color}
    />
  </svg>
)

Tooltip.defaultProps = {
  backgroundColor: colors.primary,
  color: colors.grey70,
}

const KeepCircle = ({ color }) => (
  <svg
    width="58"
    height="58"
    viewBox="0 0 58 58"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      d="M29 0.25C13.1271 0.25 0.25 13.1271 0.25 29C0.25 44.8729 13.1271 57.75 29 57.75C44.8729 57.75 57.75 44.8729 57.75 29C57.7258 13.1271 44.8487 0.25 29 0.25ZM41.0557 22.5735H39.2437L33.7595 28.9758L39.2437 35.3782H41.0557V40.5242H29.4107V35.4265H31.2227L27.6229 31.2227H26.1492V35.4265H28.2027V40.5242H16.9202V35.3782H19.3603V29V22.5735H16.9202V17.4275H19.6985V19.3361H21.1964V17.4275H23.9265V19.3361H25.4244V17.4275H28.1786V22.5011H26.125V26.7048H27.5987L31.1985 22.5011H29.3866V17.4275H41.0315V22.5735H41.0557Z"
      fill={color}
    />
  </svg>
)

KeepCircle.defaultProps = {
  color: colors.lightGrey,
}

const OK = ({ color }) => (
  <svg
    width="17"
    height="18"
    viewBox="0 0 17 18"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      d="M4.5 8.5L6.533 10.6465C6.73911 10.935 7.06732 11.1118 7.42167 11.125C7.77602 11.1382 8.11648 10.9864 8.3435 10.714L16.4375 1"
      stroke={color}
      strokeLinecap="round"
      strokeLinejoin="round"
    />
    <path
      d="M12.5872 2.3064C9.27145 0.254412 4.95279 0.941643 2.43773 3.92149C-0.0773383 6.90135 -0.0294076 11.2741 2.55037 14.1981C5.13015 17.1221 9.46283 17.7145 12.7328 15.5903C16.0028 13.4661 17.2228 9.26677 15.6 5.72115"
      stroke={color}
      strokeLinecap="round"
      strokeLinejoin="round"
    />
  </svg>
)

OK.defaultProps = {
  color: colors.black,
}

const OKBadge = ({ bgColor, color }) => (
  <svg
    width="20"
    height="21"
    viewBox="0 0 20 21"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <rect fill={bgColor} y="1" width="20" height="19" rx="9.5" />
    <path
      stroke={color}
      d="M6.96484 9.67466L8.50359 11.2987C8.6596 11.5171 8.90802 11.6508 9.17622 11.6608C9.44442 11.6708 9.70211 11.5559 9.87394 11.3498L16.0002 4"
      strokeLinecap="round"
      strokeLinejoin="round"
    />
    <path
      stroke={color}
      d="M13.0858 4.98857C10.5762 3.436 7.30743 3.95597 5.40381 6.21059C3.50019 8.46521 3.53647 11.7737 5.48907 13.9861C7.44167 16.1984 10.721 16.6467 13.196 15.0395C15.6711 13.4323 16.5945 10.2549 15.3662 7.57224"
      strokeLinecap="round"
      strokeLinejoin="round"
    />
  </svg>
)

OKBadge.defaultProps = {
  bgColor: colors.bgSuccess,
  color: colors.success,
}

const PendingBadge = ({ bgColor, color }) => (
  <svg
    width="20"
    height="21"
    viewBox="0 0 20 21"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <rect fill={bgColor} y="1" width="20" height="19" rx="9.5" />
    <svg width="14" height="14" x="3" y="3.5" fill="none">
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M7 13C10.3137 13 13 10.3137 13 7C13 3.68629 10.3137 1 7 1C3.68629 1 1 3.68629 1 7C1 10.3137 3.68629 13 7 13Z"
        stroke={color}
        strokeLinecap="round"
        strokeLinejoin="round"
      />
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M7 6.99979V4.85693V6.99979Z"
        fill="#4C4C4C"
      />
      <path
        d="M7 6.99979V4.85693"
        stroke={color}
        strokeLinecap="round"
        strokeLinejoin="round"
      />
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M7 7L9.67829 9.67886L7 7Z"
        fill="#4C4C4C"
      />
      <path
        d="M7 7L9.67829 9.67886"
        stroke={color}
        strokeLinecap="round"
        strokeLinejoin="round"
      />
    </svg>
  </svg>
)

PendingBadge.defaultProps = {
  bgColor: colors.bgPending,
  color: colors.pending,
}

export {
  Keep,
  Badge,
  Cross,
  Tooltip,
  KeepCircle,
  OK,
  OKBadge,
  PendingBadge,
  ArrowRight,
  Authorizer,
  Glossary,
  Operations,
  KeepToken,
  Rewards,
  GrantContextIcon,
  MoneyWalletOpen,
  KeepGreenOutline,
  DashedLine,
  AuthorizerFormIcon,
  BeneficiaryFormIcon,
  OperatorFormIcon,
}
