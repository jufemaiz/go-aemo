package nem12

import (
	"fmt"
	"strings"
)

const (
	// ReasonUndefined where the reason is undefined.
	ReasonUndefined = Reason(-1)
	// ReasonFreeTextDescription where the reason is free text description.
	// For use in the case that other reason code descriptions cannot be reasonably
	// utilised.
	ReasonFreeTextDescription = Reason(0)
	// ReasonMeterEquipmentChanged where the reason is meter/equipment changed.
	// Where metering installation has changed.
	ReasonMeterEquipmentChanged = Reason(1)
	// ReasonExtremeWeatherConditions where the reason is extreme weather conditions.
	// Extreme weather conditions have prevented metering data collection.
	ReasonExtremeWeatherConditions = Reason(2)
	// ReasonQuarantinedPremises where the reason is quarantined premises.
	// Site under quarantine preventing access to metering installation.
	ReasonQuarantinedPremises = Reason(3)
	// ReasonDangerousDog where the reason is dangerous dog.
	// Dog has been identified as posing an immediate hazard to metering installation
	// access.
	ReasonDangerousDog = Reason(4)
	// ReasonBlankScreen where the reason is blank screen.
	// Electronic meter has blank display, could be powered off or faulty display
	// but unable to determine.
	ReasonBlankScreen = Reason(5)
	// ReasonDeEnergisedPremises where the reason is de-energised premises.
	// Blank screen on an electronic meter where the meter reader can determine that
	// the Site is de-energised or an Interval Metered Site where the MDP is providing
	// substituted metering data for a Site that is de-energised but Datastreams are
	// left active.
	ReasonDeEnergisedPremises = Reason(6)
	// ReasonUnableToLocateMeter where the reason is unable to locate meter.
	// The Site was found, but unable to locate the metering installation.
	ReasonUnableToLocateMeter = Reason(7)
	// ReasonVacantPremises where the reason is vacant premises. Meter reader believes
	// the Site is vacant.
	ReasonVacantPremises = Reason(8)
	// ReasonUnderInvestigation where the reason is under investigation. An issue
	// with the metering installation has been identified and is under investigation.
	ReasonUnderInvestigation = Reason(9)
	// ReasonLockDamagedUnableToOpen where the reason is lock damaged unable to open.
	// Unable to open lock due to damage and the lock is preventing access to the
	// metering installation.
	ReasonLockDamagedUnableToOpen = Reason(10)
	// ReasonInWrongRoute where the reason is in wrong route. Unable to obtain reading
	// due to the metering installation being in the wrong route.
	ReasonInWrongRoute = Reason(11)
	// ReasonLockedPremises where the reason is locked premises. Unable to obtain
	// access to metering installation due to Site being locked.
	ReasonLockedPremises = Reason(12)
	// ReasonLockedGate where the reason is locked gate. Locked gate at Site is
	// preventing access to metering installation.
	ReasonLockedGate = Reason(13)
	// ReasonLockedMeterBox where the reason is locked meter box. Locked meter box
	// is preventing access to metering installation.
	ReasonLockedMeterBox = Reason(14)
	// ReasonOvergrownVegetation where the reason is overgrown vegetation. Overgrown
	// vegetation at Site is preventing access to metering installation.
	ReasonOvergrownVegetation = Reason(15)
	// ReasonNoxiousWeeds where the reason is noxious weeds. Noxious weeds at Site
	// are preventing access to metering installation.
	ReasonNoxiousWeeds = Reason(16)
	// ReasonUnsafeEquipmentLocation where the reason is unsafe equipment/location.
	// The equipment or the location of the metering installation has been identified
	// as unsafe (other than meter being high).
	ReasonUnsafeEquipmentLocation = Reason(17)
	// ReasonReadLessThanPrevious where the reason is read less than previous.
	// Current Meter Reading obtained is less than previous Meter Reading, no evidence
	// of tampering and no reverse energy observed.
	ReasonReadLessThanPrevious = Reason(18)
	// ReasonConsumerWanted where the reason is consumer wantedvconsumer
	// wanted.
	ReasonConsumerWanted = Reason(19)
	// ReasonDamagedEquipmentPanel where the reason is damaged equipment/panel.
	// The equipment or the panel of the metering installation has been damaged
	// but has not been identified as unsafe.
	ReasonDamagedEquipmentPanel = Reason(20)
	// ReasonMainSwitchOff where the reason is main switch off. Blank screen on an
	// electronic meter where the meter reader can determine that the main switch
	// has been turned off or Interval Metered Site where the MDP is providing
	// substituted metering data for a Site that the main switch is off but
	// Datastreams are left active.
	ReasonMainSwitchOff = Reason(21)
	// ReasonMeterEquipmentSealsMissing where the reason is meter/equipment seals
	// missing. One or more seals are missing from the metering installation, no
	// tampering has been identified.
	ReasonMeterEquipmentSealsMissing = Reason(22)
	// ReasonReaderError where the reason is reader error. MDP identified that
	// Meter Reading provided by the meter reader was incorrect.
	ReasonReaderError = Reason(23)
	// ReasonSubstitutedReplacedDataDataCorrection where the reason is
	// substituted/replaced data (data correction). Interval Meter Reading – MDP
	// replaced erroneous data for specific Intervals.
	ReasonSubstitutedReplacedDataDataCorrection = Reason(24)
	// ReasonUnableToLocatePremises where the reason is unable to locate premises.
	// Unable to locate Site.
	ReasonUnableToLocatePremises = Reason(25)
	// ReasonNegativeConsumptionGeneration where the reason is negative
	// consumption (generation). Accumulation Meter where the previous Meter
	// Reading is higher than the current Meter Reading, generally Site will have
	// generation.
	ReasonNegativeConsumptionGeneration = Reason(26)
	// ReasonRolr where the reason is rolr. To be used when transferring End User
	// as a result of a RoLR Event.
	ReasonRolr = Reason(27)
	// ReasonCtVtFault where the reason is ct/vt fault. MDP has corrected data due
	// to a known instrument transformer (CT/VT) fault.
	ReasonCtVtFault = Reason(28)
	// ReasonRelayFaultyDamaged where the reason is relay faulty/damaged. Meter
	// reader has identified the relay device within the metering installation is
	// faulty.
	ReasonRelayFaultyDamaged = Reason(29)
	// ReasonMeterStopSwitchOn where the reason is meter stop switch on.
	// (DEPRECATED).
	ReasonMeterStopSwitchOn = Reason(30)
	// ReasonNotAllMetersRead where the reason is not all meters read. Readings
	// for all meters linked to the Site have not been received by the MDP
	// (typically as a result of a non-Scheduled Meter Reading).
	ReasonNotAllMetersRead = Reason(31)
	// ReasonReEnergisedWithoutReadings where the reason is re-energised without
	// readings. Unable to obtain Meter Readings due to exceptional circumstances
	// when the Site is re-energised outside of standard practice.
	ReasonReEnergisedWithoutReadings = Reason(32)
	// ReasonDeEnergisedWithoutReadings where the reason is de-energised without
	// readings. Unable to obtain Meter Readings at the time of de-energisation
	// including disconnection for non-payment.
	ReasonDeEnergisedWithoutReadings = Reason(33)
	// ReasonMeterNotInHandheld where the reason is meter not in handheld.
	// Unexpected meter found on Site (new meter or additional meter).
	ReasonMeterNotInHandheld = Reason(34)
	// ReasonTimeswitchFaultyResetRequired where the reason is timeswitch
	// faulty/reset required. Meter reader has identified the time switching
	// device within the metering installation is faulty and required resetting.
	ReasonTimeswitchFaultyResetRequired = Reason(35)
	// ReasonMeterHighLadderRequired where the reason is meter high/ladder
	// required. Meter in a high position requiring a ladder to obtain Meter
	// Reading.
	ReasonMeterHighLadderRequired = Reason(36)
	// ReasonMeterUnderChurn where the reason is meter under churn. MDP has
	// Substituted data based on metering data not being received from relevant
	// MDP.
	ReasonMeterUnderChurn = Reason(37)
	// ReasonUnmarriedLock where the reason is unmarried lock. Site has two or
	// more locks, one of which is a power industry lock and they have not been
	// interlocked together correctly to allow access to the Site.
	ReasonUnmarriedLock = Reason(38)
	// ReasonReverseEnergyObserved where the reason is reverse energy observed.
	// Reverse energy observed where Site isn’t expected to have reverse energy.
	ReasonReverseEnergyObserved = Reason(39)
	// ReasonUnrestrainedLivestock where the reason is unrestrained livestock.
	// Meter reader observed that livestock is roaming free on Site and could
	// potentially be hazardous, or access wasn’t obtained due to potential for
	// livestock to escape.
	ReasonUnrestrainedLivestock = Reason(40)
	// ReasonFaultyMeterDisplayDials where the reason is faulty meter
	// display/dials. Display or dials on the meter are faulty and Site is not
	// de-energised nor is the display blank on an electronic meter.
	ReasonFaultyMeterDisplayDials = Reason(41)
	// ReasonChannelAddedRemoved where the reason is channel added/removed. MDP
	// obtained metering data for a channel that has been added or substituted
	// metering data where a channel has been removed but the Datastream is still
	// active in MSATS.
	ReasonChannelAddedRemoved = Reason(42)
	// ReasonPowerOutage where the reason is power outage. Interval Meter –
	// Metering data for Intervals have been Substituted due to power not being
	// available at the metering installation.
	ReasonPowerOutage = Reason(43)
	// ReasonMeterTesting where the reason is meter testing. MDP identifies meter
	// has been under testing regime and has provided substituted metering data to
	// reflect energy consumption during testing period.
	ReasonMeterTesting = Reason(44)
	// ReasonReadingsFailedToValidate where the reason is readings failed to
	// validate. Meter Readings have been loaded into MDP’s system, but have
	// failed Validation and have been Substituted.
	ReasonReadingsFailedToValidate = Reason(45)
	// ReasonExtremeWeatherHot where the reason is extreme weather/hot.
	ReasonExtremeWeatherHot = Reason(46)
	// ReasonRefusedAccess where the reason is refused access. The End User
	// refused to provide access when requested.
	ReasonRefusedAccess = Reason(47)
	// ReasonDogOnPremises where the reason is dog on premises. Meter reader has
	// identified that there is a dog on the Site but has been unable to determine
	// if the dog is dangerous.
	ReasonDogOnPremises = Reason(48)
	// ReasonWetPaint where the reason is wet paint (DEPRECATED).
	ReasonWetPaint = Reason(49)
	// ReasonWrongTarif where the reason is wrong tarif (DEPRECATED).
	ReasonWrongTarif = Reason(50)
	// ReasonInstallationDemolished where the reason is installation demolished.
	// Metering installation no longer exists at the Site.
	ReasonInstallationDemolished = Reason(51)
	// ReasonAccessBlocked where the reason is access – blocked. Used when there
	// are items blocking safe access to the meter or Site.
	ReasonAccessBlocked = Reason(52)
	// ReasonPestsInMeterBox where the reason is pests in meter box. Pests have
	// been identified within the meter box that poses a risk to metering data
	// accuracy, safety of the metering installation or a hazard to the meter
	// reader.
	ReasonPestsInMeterBox = Reason(53)
	// ReasonMeterBoxDamagedFaulty where the reason is meter box damaged/faulty.
	// Meter reader identifies that the meter box is damaged or faulty and the
	// mechanical protection or weather proofing of the metering installation is
	// compromised as a result.
	ReasonMeterBoxDamagedFaulty = Reason(54)
	// ReasonDialsObscured where the reason is dials obscured. Meter reader unable
	// to obtain Meter Reading due to meter dials being obscured, meter face
	// painted over, viewing panel in locked meter box with pvc panel misted
	// over/faded/mouldy etc. No evidence of tampering.
	ReasonDialsObscured = Reason(55)
	// ReasonMeterOkSupplyFailure where the reason is meter ok – supply failure.
	// (DEPRECATED).
	ReasonMeterOkSupplyFailure = Reason(58)
	// ReasonIllegalConnection where the reason is illegal connection. Meter
	// reader has identified that the Site has been illegally connected.
	ReasonIllegalConnection = Reason(60)
	// ReasonEquipmentTampered where the reason is equipment tampered. Meter
	// reader identified that the metering installation has been tampered with and
	// the recording of energy consumption may have been affected as a result.
	ReasonEquipmentTampered = Reason(61)
	// ReasonNsrdWindowExpired where the reason is nsrd window expired. Where the
	// NSRD window has expired and the meter reader has been unable to deliver
	// Actual Meter Readings.
	ReasonNsrdWindowExpired = Reason(62)
	// ReasonKeyRequired where the reason is key required. Meter reader typically
	// has access to the key but was unable to obtain/locate the key at the time
	// of Meter Reading.
	ReasonKeyRequired = Reason(64)
	// ReasonWrongKeyProvided where the reason is wrong key provided. Meter reader
	// has been provided with a key but the key no longer opens the lock.
	ReasonWrongKeyProvided = Reason(65)
	// ReasonZeroConsumption where the reason is zero consumption. Where a Site
	// has known zero consumption and the Site is not deenergised in MSATS but no
	// energy is flowing to the meter.
	ReasonZeroConsumption = Reason(68)
	// ReasonReadingExceedsSubstitute where the reason is reading exceeds
	// substitute. Re-Substituted data that has been modified to improve the
	// smoothing of energy to align with the next Actual Meter Reading.
	ReasonReadingExceedsSubstitute = Reason(69)
	// ReasonProbeReportsTampering where the reason is probe reports tampering.
	// (DEPRECATED).
	ReasonProbeReportsTampering = Reason(70)
	// ReasonProbeReadError where the reason is probe read error. Data collector
	// unable to collect the metering data due to the meter probe being unable to
	// extract the metering data.
	ReasonProbeReadError = Reason(71)
	// ReasonReCalculatedBasedOnActualMeteringData where the reason is
	// re-calculated based on actual metering data. MDP received Actual Meter
	// Readings and prior Substitutes have been amended.
	ReasonReCalculatedBasedOnActualMeteringData = Reason(72)
	// ReasonLowConsumption where the reason is low consumption. Meter Reading
	// failed Validation as being too low based on Historical Data and has been
	// either left as an actual or replaced by a Substitute.
	ReasonLowConsumption = Reason(73)
	// ReasonHighConsumption where the reason is high consumption. Meter Reading
	// failed Validation as being too high based on Historical Data and has been
	// either left as an actual or replaced by a Substitute.
	ReasonHighConsumption = Reason(74)
	// ReasonCustomerRead where the reason is customer read. Meter Reading
	// provided to the MDP by the End User. (Only applicable in Jurisdictions
	// where End User Meter Readings are allowed).
	ReasonCustomerRead = Reason(75)
	// ReasonCommunicationsFault where the reason is communications fault. Meter
	// reader attempted to read the meter but was unable due to not being able to
	// remotely communicate with the meter.
	ReasonCommunicationsFault = Reason(76)
	// ReasonEstimationForecast where the reason is estimation forecast. Optional
	// reason code that can be applied to Estimations.
	ReasonEstimationForecast = Reason(77)
	// ReasonNullData where the reason is null data. For Interval Meters where no
	// metering data was received and Substitutes created to cover this period.
	ReasonNullData = Reason(78)
	// ReasonPowerOutageAlarm where the reason is power outage alarm. For Interval
	// Meters where a power outage has been detected by the meter.
	ReasonPowerOutageAlarm = Reason(79)
	// ReasonShortIntervalAlarm where the reason is short interval alarm. For
	// Interval Meters where the time in the meter is slow and has now been
	// corrected, resulting in the interval metering data not being a full 15 or
	// 30 minutes in length.
	ReasonShortIntervalAlarm = Reason(80)
	// ReasonLongIntervalAlarm where the reason is long interval alarm. For
	// Interval Meters where the time in the meter is fast and has now been
	// corrected, resulting in the interval metering data exceeding a full 15 or
	// 30 minutes in length.
	ReasonLongIntervalAlarm = Reason(81)
	// ReasonCrcError where the reason is crc error (DEPRECATED).
	ReasonCrcError = Reason(82)
	// ReasonRAMChecksumError where the reason is ram checksum error (DEPRECATED).
	ReasonRAMChecksumError = Reason(83)
	// ReasonROMChecksumError where the reason is rom checksum error (DEPRECATED).
	ReasonROMChecksumError = Reason(84)
	// ReasonDataMissingAlarm where the reason is data missing alarm (DEPRECATED).
	ReasonDataMissingAlarm = Reason(85)
	// ReasonClockErrorAlarm where the reason is clock error alarm (DEPRECATED).
	ReasonClockErrorAlarm = Reason(86)
	// ReasonResetOccurred where the reason is reset occurred. Resetting of the
	// meter due to re-programming, change of configuration or firmware upgrade
	// etc.
	ReasonResetOccurred = Reason(87)
	// ReasonWatchdogTimeoutAlarm where the reason is watchdog timeout alarm.
	// (DEPRECATED).
	ReasonWatchdogTimeoutAlarm = Reason(88)
	// ReasonTimeResetOccurred where the reason is time reset occurred. Where a
	// time reset has occurred within the metering installation.
	ReasonTimeResetOccurred = Reason(89)
	// ReasonTestMode where the reason is test mode (DEPRECATED).
	ReasonTestMode = Reason(90)
	// ReasonLoadControl where the reason is load control (DEPRECATED).
	ReasonLoadControl = Reason(91)
	// ReasonAddedIntervalDataCorrection where the reason is added interval (data
	// correction) (DEPRECATED).
	ReasonAddedIntervalDataCorrection = Reason(92)
	// ReasonReplacedIntervalDataCorrection where the reason is replaced interval
	// (data correction) (DEPRECATED).
	ReasonReplacedIntervalDataCorrection = Reason(93)
	// ReasonEstimatedIntervalDataCorrection where the reason is estimated
	// interval (data correction) (DEPRECATED).
	ReasonEstimatedIntervalDataCorrection = Reason(94)
	// ReasonPulseOverflowAlarm where the reason is pulse overflow alarm.
	// (DEPRECATED).
	ReasonPulseOverflowAlarm = Reason(95)
	// ReasonDataOutOfLimits where the reason is data out of limits (DEPRECATED).
	ReasonDataOutOfLimits = Reason(96)
	// ReasonExcludedData where the reason is excluded data (DEPRECATED).
	ReasonExcludedData = Reason(97)
	// ReasonParityError where the reason is parity error (DEPRECATED).
	ReasonParityError = Reason(98)
	// ReasonEnergyTypeRegisterChanged where the reason is energy type (register
	// changed) (DEPRECATED).
	ReasonEnergyTypeRegisterChanged = Reason(99)
)

var (
	reasons = []Reason{ //nolint:gochecknoglobals
		ReasonFreeTextDescription,
		ReasonMeterEquipmentChanged,
		ReasonExtremeWeatherConditions,
		ReasonQuarantinedPremises,
		ReasonDangerousDog,
		ReasonBlankScreen,
		ReasonDeEnergisedPremises,
		ReasonUnableToLocateMeter,
		ReasonVacantPremises,
		ReasonUnderInvestigation,
		ReasonLockDamagedUnableToOpen,
		ReasonInWrongRoute,
		ReasonLockedPremises,
		ReasonLockedGate,
		ReasonLockedMeterBox,
		ReasonOvergrownVegetation,
		ReasonNoxiousWeeds,
		ReasonUnsafeEquipmentLocation,
		ReasonReadLessThanPrevious,
		ReasonConsumerWanted,
		ReasonDamagedEquipmentPanel,
		ReasonMainSwitchOff,
		ReasonMeterEquipmentSealsMissing,
		ReasonReaderError,
		ReasonSubstitutedReplacedDataDataCorrection,
		ReasonUnableToLocatePremises,
		ReasonNegativeConsumptionGeneration,
		ReasonRolr,
		ReasonCtVtFault,
		ReasonRelayFaultyDamaged,
		ReasonMeterStopSwitchOn,
		ReasonNotAllMetersRead,
		ReasonReEnergisedWithoutReadings,
		ReasonDeEnergisedWithoutReadings,
		ReasonMeterNotInHandheld,
		ReasonTimeswitchFaultyResetRequired,
		ReasonMeterHighLadderRequired,
		ReasonMeterUnderChurn,
		ReasonUnmarriedLock,
		ReasonReverseEnergyObserved,
		ReasonUnrestrainedLivestock,
		ReasonFaultyMeterDisplayDials,
		ReasonChannelAddedRemoved,
		ReasonPowerOutage,
		ReasonMeterTesting,
		ReasonReadingsFailedToValidate,
		ReasonExtremeWeatherHot,
		ReasonRefusedAccess,
		ReasonDogOnPremises,
		ReasonWetPaint,
		ReasonWrongTarif,
		ReasonInstallationDemolished,
		ReasonAccessBlocked,
		ReasonPestsInMeterBox,
		ReasonMeterBoxDamagedFaulty,
		ReasonDialsObscured,
		ReasonMeterOkSupplyFailure,
		ReasonIllegalConnection,
		ReasonEquipmentTampered,
		ReasonNsrdWindowExpired,
		ReasonKeyRequired,
		ReasonWrongKeyProvided,
		ReasonZeroConsumption,
		ReasonReadingExceedsSubstitute,
		ReasonProbeReportsTampering,
		ReasonProbeReadError,
		ReasonReCalculatedBasedOnActualMeteringData,
		ReasonLowConsumption,
		ReasonHighConsumption,
		ReasonCustomerRead,
		ReasonCommunicationsFault,
		ReasonEstimationForecast,
		ReasonNullData,
		ReasonPowerOutageAlarm,
		ReasonShortIntervalAlarm,
		ReasonLongIntervalAlarm,
		ReasonCrcError,
		ReasonRAMChecksumError,
		ReasonROMChecksumError,
		ReasonDataMissingAlarm,
		ReasonClockErrorAlarm,
		ReasonResetOccurred,
		ReasonWatchdogTimeoutAlarm,
		ReasonTimeResetOccurred,
		ReasonTestMode,
		ReasonLoadControl,
		ReasonAddedIntervalDataCorrection,
		ReasonReplacedIntervalDataCorrection,
		ReasonEstimatedIntervalDataCorrection,
		ReasonPulseOverflowAlarm,
		ReasonDataOutOfLimits,
		ReasonExcludedData,
		ReasonParityError,
		ReasonEnergyTypeRegisterChanged,
	}

	// ReasonName a mapping of the reason to the string of the reason.
	ReasonName = map[Reason]string{ //nolint:dupl,gochecknoglobals
		ReasonFreeTextDescription:                   "0",
		ReasonMeterEquipmentChanged:                 "1",
		ReasonExtremeWeatherConditions:              "2",
		ReasonQuarantinedPremises:                   "3",
		ReasonDangerousDog:                          "4",
		ReasonBlankScreen:                           "5",
		ReasonDeEnergisedPremises:                   "6",
		ReasonUnableToLocateMeter:                   "7",
		ReasonVacantPremises:                        "8",
		ReasonUnderInvestigation:                    "9",
		ReasonLockDamagedUnableToOpen:               "10",
		ReasonInWrongRoute:                          "11",
		ReasonLockedPremises:                        "12",
		ReasonLockedGate:                            "13",
		ReasonLockedMeterBox:                        "14",
		ReasonOvergrownVegetation:                   "15",
		ReasonNoxiousWeeds:                          "16",
		ReasonUnsafeEquipmentLocation:               "17",
		ReasonReadLessThanPrevious:                  "18",
		ReasonConsumerWanted:                        "19",
		ReasonDamagedEquipmentPanel:                 "20",
		ReasonMainSwitchOff:                         "21",
		ReasonMeterEquipmentSealsMissing:            "22",
		ReasonReaderError:                           "23",
		ReasonSubstitutedReplacedDataDataCorrection: "24",
		ReasonUnableToLocatePremises:                "25",
		ReasonNegativeConsumptionGeneration:         "26",
		ReasonRolr:                                  "27",
		ReasonCtVtFault:                             "28",
		ReasonRelayFaultyDamaged:                    "29",
		ReasonMeterStopSwitchOn:                     "30",
		ReasonNotAllMetersRead:                      "31",
		ReasonReEnergisedWithoutReadings:            "32",
		ReasonDeEnergisedWithoutReadings:            "33",
		ReasonMeterNotInHandheld:                    "34",
		ReasonTimeswitchFaultyResetRequired:         "35",
		ReasonMeterHighLadderRequired:               "36",
		ReasonMeterUnderChurn:                       "37",
		ReasonUnmarriedLock:                         "38",
		ReasonReverseEnergyObserved:                 "39",
		ReasonUnrestrainedLivestock:                 "40",
		ReasonFaultyMeterDisplayDials:               "41",
		ReasonChannelAddedRemoved:                   "42",
		ReasonPowerOutage:                           "43",
		ReasonMeterTesting:                          "44",
		ReasonReadingsFailedToValidate:              "45",
		ReasonExtremeWeatherHot:                     "46",
		ReasonRefusedAccess:                         "47",
		ReasonDogOnPremises:                         "48",
		ReasonWetPaint:                              "49",
		ReasonWrongTarif:                            "50",
		ReasonInstallationDemolished:                "51",
		ReasonAccessBlocked:                         "52",
		ReasonPestsInMeterBox:                       "53",
		ReasonMeterBoxDamagedFaulty:                 "54",
		ReasonDialsObscured:                         "55",
		ReasonMeterOkSupplyFailure:                  "58",
		ReasonIllegalConnection:                     "60",
		ReasonEquipmentTampered:                     "61",
		ReasonNsrdWindowExpired:                     "62",
		ReasonKeyRequired:                           "64",
		ReasonWrongKeyProvided:                      "65",
		ReasonZeroConsumption:                       "68",
		ReasonReadingExceedsSubstitute:              "69",
		ReasonProbeReportsTampering:                 "70",
		ReasonProbeReadError:                        "71",
		ReasonReCalculatedBasedOnActualMeteringData: "72",
		ReasonLowConsumption:                        "73",
		ReasonHighConsumption:                       "74",
		ReasonCustomerRead:                          "75",
		ReasonCommunicationsFault:                   "76",
		ReasonEstimationForecast:                    "77",
		ReasonNullData:                              "78",
		ReasonPowerOutageAlarm:                      "79",
		ReasonShortIntervalAlarm:                    "80",
		ReasonLongIntervalAlarm:                     "81",
		ReasonCrcError:                              "82",
		ReasonRAMChecksumError:                      "83",
		ReasonROMChecksumError:                      "84",
		ReasonDataMissingAlarm:                      "85",
		ReasonClockErrorAlarm:                       "86",
		ReasonResetOccurred:                         "87",
		ReasonWatchdogTimeoutAlarm:                  "88",
		ReasonTimeResetOccurred:                     "89",
		ReasonTestMode:                              "90",
		ReasonLoadControl:                           "91",
		ReasonAddedIntervalDataCorrection:           "92",
		ReasonReplacedIntervalDataCorrection:        "93",
		ReasonEstimatedIntervalDataCorrection:       "94",
		ReasonPulseOverflowAlarm:                    "95",
		ReasonDataOutOfLimits:                       "96",
		ReasonExcludedData:                          "97",
		ReasonParityError:                           "98",
		ReasonEnergyTypeRegisterChanged:             "99",
	}

	// ReasonValue a mapping of the string of the reason code to the reason.
	ReasonValue = map[string]Reason{ //nolint:gochecknoglobals
		"0":  ReasonFreeTextDescription,
		"1":  ReasonMeterEquipmentChanged,
		"2":  ReasonExtremeWeatherConditions,
		"3":  ReasonQuarantinedPremises,
		"4":  ReasonDangerousDog,
		"5":  ReasonBlankScreen,
		"6":  ReasonDeEnergisedPremises,
		"7":  ReasonUnableToLocateMeter,
		"8":  ReasonVacantPremises,
		"9":  ReasonUnderInvestigation,
		"10": ReasonLockDamagedUnableToOpen,
		"11": ReasonInWrongRoute,
		"12": ReasonLockedPremises,
		"13": ReasonLockedGate,
		"14": ReasonLockedMeterBox,
		"15": ReasonOvergrownVegetation,
		"16": ReasonNoxiousWeeds,
		"17": ReasonUnsafeEquipmentLocation,
		"18": ReasonReadLessThanPrevious,
		"19": ReasonConsumerWanted,
		"20": ReasonDamagedEquipmentPanel,
		"21": ReasonMainSwitchOff,
		"22": ReasonMeterEquipmentSealsMissing,
		"23": ReasonReaderError,
		"24": ReasonSubstitutedReplacedDataDataCorrection,
		"25": ReasonUnableToLocatePremises,
		"26": ReasonNegativeConsumptionGeneration,
		"27": ReasonRolr,
		"28": ReasonCtVtFault,
		"29": ReasonRelayFaultyDamaged,
		"30": ReasonMeterStopSwitchOn,
		"31": ReasonNotAllMetersRead,
		"32": ReasonReEnergisedWithoutReadings,
		"33": ReasonDeEnergisedWithoutReadings,
		"34": ReasonMeterNotInHandheld,
		"35": ReasonTimeswitchFaultyResetRequired,
		"36": ReasonMeterHighLadderRequired,
		"37": ReasonMeterUnderChurn,
		"38": ReasonUnmarriedLock,
		"39": ReasonReverseEnergyObserved,
		"40": ReasonUnrestrainedLivestock,
		"41": ReasonFaultyMeterDisplayDials,
		"42": ReasonChannelAddedRemoved,
		"43": ReasonPowerOutage,
		"44": ReasonMeterTesting,
		"45": ReasonReadingsFailedToValidate,
		"46": ReasonExtremeWeatherHot,
		"47": ReasonRefusedAccess,
		"48": ReasonDogOnPremises,
		"49": ReasonWetPaint,
		"50": ReasonWrongTarif,
		"51": ReasonInstallationDemolished,
		"52": ReasonAccessBlocked,
		"53": ReasonPestsInMeterBox,
		"54": ReasonMeterBoxDamagedFaulty,
		"55": ReasonDialsObscured,
		"58": ReasonMeterOkSupplyFailure,
		"60": ReasonIllegalConnection,
		"61": ReasonEquipmentTampered,
		"62": ReasonNsrdWindowExpired,
		"64": ReasonKeyRequired,
		"65": ReasonWrongKeyProvided,
		"68": ReasonZeroConsumption,
		"69": ReasonReadingExceedsSubstitute,
		"70": ReasonProbeReportsTampering,
		"71": ReasonProbeReadError,
		"72": ReasonReCalculatedBasedOnActualMeteringData,
		"73": ReasonLowConsumption,
		"74": ReasonHighConsumption,
		"75": ReasonCustomerRead,
		"76": ReasonCommunicationsFault,
		"77": ReasonEstimationForecast,
		"78": ReasonNullData,
		"79": ReasonPowerOutageAlarm,
		"80": ReasonShortIntervalAlarm,
		"81": ReasonLongIntervalAlarm,
		"82": ReasonCrcError,
		"83": ReasonRAMChecksumError,
		"84": ReasonROMChecksumError,
		"85": ReasonDataMissingAlarm,
		"86": ReasonClockErrorAlarm,
		"87": ReasonResetOccurred,
		"88": ReasonWatchdogTimeoutAlarm,
		"89": ReasonTimeResetOccurred,
		"90": ReasonTestMode,
		"91": ReasonLoadControl,
		"92": ReasonAddedIntervalDataCorrection,
		"93": ReasonReplacedIntervalDataCorrection,
		"94": ReasonEstimatedIntervalDataCorrection,
		"95": ReasonPulseOverflowAlarm,
		"96": ReasonDataOutOfLimits,
		"97": ReasonExcludedData,
		"98": ReasonParityError,
		"99": ReasonEnergyTypeRegisterChanged,
	}

	// reasonDescriptions maps reasons to descriptions.
	reasonDescriptions = map[Reason]string{ //nolint:dupl,gochecknoglobals
		ReasonFreeTextDescription:                   "Free text description",
		ReasonMeterEquipmentChanged:                 "Meter/equipment changed",
		ReasonExtremeWeatherConditions:              "Extreme weather conditions",
		ReasonQuarantinedPremises:                   "Quarantined premises",
		ReasonDangerousDog:                          "Dangerous dog",
		ReasonBlankScreen:                           "Blank screen",
		ReasonDeEnergisedPremises:                   "De-energised premises",
		ReasonUnableToLocateMeter:                   "Unable to locate meter",
		ReasonVacantPremises:                        "Vacant premises",
		ReasonUnderInvestigation:                    "Under investigation",
		ReasonLockDamagedUnableToOpen:               "Lock damaged unable to open",
		ReasonInWrongRoute:                          "In wrong route",
		ReasonLockedPremises:                        "Locked premises",
		ReasonLockedGate:                            "Locked gate",
		ReasonLockedMeterBox:                        "Locked meter box",
		ReasonOvergrownVegetation:                   "Overgrown vegetation",
		ReasonNoxiousWeeds:                          "Noxious weeds",
		ReasonUnsafeEquipmentLocation:               "Unsafe equipment/location",
		ReasonReadLessThanPrevious:                  "Read less than previous",
		ReasonConsumerWanted:                        "Consumer wanted.",
		ReasonDamagedEquipmentPanel:                 "Damaged equipment/panel",
		ReasonMainSwitchOff:                         "Main switch off",
		ReasonMeterEquipmentSealsMissing:            "Meter/equipment seals missing",
		ReasonReaderError:                           "Reader error",
		ReasonSubstitutedReplacedDataDataCorrection: "Substituted/replaced data (data correction)",
		ReasonUnableToLocatePremises:                "Unable to locate premises",
		ReasonNegativeConsumptionGeneration:         "Negative consumption (generation)",
		ReasonRolr:                                  "ROLR",
		ReasonCtVtFault:                             "CT/VT fault",
		ReasonRelayFaultyDamaged:                    "Relay faulty/damaged",
		ReasonMeterStopSwitchOn:                     "Meter stop switch on",
		ReasonNotAllMetersRead:                      "Not all meters read",
		ReasonReEnergisedWithoutReadings:            "Re-energised without readings",
		ReasonDeEnergisedWithoutReadings:            "De-energised without readings",
		ReasonMeterNotInHandheld:                    "Meter not in handheld",
		ReasonTimeswitchFaultyResetRequired:         "Timeswitch faulty/reset required",
		ReasonMeterHighLadderRequired:               "Meter high/ladder required",
		ReasonMeterUnderChurn:                       "Meter under churn",
		ReasonUnmarriedLock:                         "Unmarried lock",
		ReasonReverseEnergyObserved:                 "Reverse energy observed",
		ReasonUnrestrainedLivestock:                 "Unrestrained livestock",
		ReasonFaultyMeterDisplayDials:               "Faulty Meter display/dials",
		ReasonChannelAddedRemoved:                   "Channel added/removed",
		ReasonPowerOutage:                           "Power outage",
		ReasonMeterTesting:                          "Meter testing",
		ReasonReadingsFailedToValidate:              "Readings failed to validate",
		ReasonExtremeWeatherHot:                     "Extreme weather/hot.",
		ReasonRefusedAccess:                         "Refused access",
		ReasonDogOnPremises:                         "Dog on premises",
		ReasonWetPaint:                              "Wet paint",
		ReasonWrongTarif:                            "Wrong tarif",
		ReasonInstallationDemolished:                "Installation demolished",
		ReasonAccessBlocked:                         "Access – blocked",
		ReasonPestsInMeterBox:                       "Pests in meter box",
		ReasonMeterBoxDamagedFaulty:                 "Meter box damaged/faulty",
		ReasonDialsObscured:                         "Dials obscured",
		ReasonMeterOkSupplyFailure:                  "Meter ok – supply failure",
		ReasonIllegalConnection:                     "Illegal connection",
		ReasonEquipmentTampered:                     "Equipment tampered",
		ReasonNsrdWindowExpired:                     "NSRD window expired",
		ReasonKeyRequired:                           "Key required",
		ReasonWrongKeyProvided:                      "Wrong key provided",
		ReasonZeroConsumption:                       "Zero consumption",
		ReasonReadingExceedsSubstitute:              "Reading exceeds Substitute",
		ReasonProbeReportsTampering:                 "Probe reports tampering",
		ReasonProbeReadError:                        "Probe read error",
		ReasonReCalculatedBasedOnActualMeteringData: "Re-calculated based on Actual Metering Data",
		ReasonLowConsumption:                        "Low consumption",
		ReasonHighConsumption:                       "High consumption",
		ReasonCustomerRead:                          "Customer read",
		ReasonCommunicationsFault:                   "Communications fault",
		ReasonEstimationForecast:                    "Estimation Forecast",
		ReasonNullData:                              "Null Data",
		ReasonPowerOutageAlarm:                      "Power Outage Alarm",
		ReasonShortIntervalAlarm:                    "Short Interval Alarm",
		ReasonLongIntervalAlarm:                     "Long Interval Alarm",
		ReasonCrcError:                              "CRC error",
		ReasonRAMChecksumError:                      "RAM checksum error",
		ReasonROMChecksumError:                      "ROM checksum error",
		ReasonDataMissingAlarm:                      "Data missing alarm",
		ReasonClockErrorAlarm:                       "Clock error alarm",
		ReasonResetOccurred:                         "Reset occurred",
		ReasonWatchdogTimeoutAlarm:                  "Watchdog timeout alarm",
		ReasonTimeResetOccurred:                     "Time reset occurred",
		ReasonTestMode:                              "Test mode",
		ReasonLoadControl:                           "Load control",
		ReasonAddedIntervalDataCorrection:           "Added interval (data correction)",
		ReasonReplacedIntervalDataCorrection:        "Replaced interval (data correction)",
		ReasonEstimatedIntervalDataCorrection:       "Estimated interval (data correction)",
		ReasonPulseOverflowAlarm:                    "Pulse overflow alarm",
		ReasonDataOutOfLimits:                       "Data out of limits",
		ReasonExcludedData:                          "Excluded data",
		ReasonParityError:                           "Parity error",
		ReasonEnergyTypeRegisterChanged:             "Energy type (register changed)",
	}

	reasonsDeprecated = map[Reason]bool{ //nolint:gochecknoglobals
		ReasonConsumerWanted:                  true,
		ReasonMeterStopSwitchOn:               true,
		ReasonExtremeWeatherHot:               true,
		ReasonWetPaint:                        true,
		ReasonWrongTarif:                      true,
		ReasonMeterOkSupplyFailure:            true,
		ReasonProbeReportsTampering:           true,
		ReasonCrcError:                        true,
		ReasonRAMChecksumError:                true,
		ReasonROMChecksumError:                true,
		ReasonDataMissingAlarm:                true,
		ReasonClockErrorAlarm:                 true,
		ReasonWatchdogTimeoutAlarm:            true,
		ReasonTestMode:                        true,
		ReasonLoadControl:                     true,
		ReasonAddedIntervalDataCorrection:     true,
		ReasonReplacedIntervalDataCorrection:  true,
		ReasonEstimatedIntervalDataCorrection: true,
		ReasonPulseOverflowAlarm:              true,
		ReasonDataOutOfLimits:                 true,
		ReasonExcludedData:                    true,
		ReasonParityError:                     true,
		ReasonEnergyTypeRegisterChanged:       true,
	}
)

// A Reason represents the value of the reason code field of a NEM12 record.
type Reason int

// Reasons returns all reasons.
func Reasons() []Reason {
	return reasons
}

// DeprecatedReasons returns all reasons that are deprecated.
func DeprecatedReasons() []Reason {
	resp := []Reason{}

	for _, r := range reasons {
		_, ok := reasonsDeprecated[r]
		if ok {
			resp = append(resp, r)
		}
	}

	return resp
}

// ActiveReasons returns all reasons that are not deprecated.
func ActiveReasons() []Reason {
	resp := []Reason{}

	for _, r := range reasons {
		_, ok := reasonsDeprecated[r]
		if !ok {
			resp = append(resp, r)
		}
	}

	return resp
}

// NewReason returns a new reason, along with errors if not valid.
func NewReason(s string) (Reason, error) {
	if s == "" {
		return ReasonUndefined, ErrReasonCodeNil
	}

	r, ok := ReasonValue[strings.ToUpper(s)]
	if !ok {
		return ReasonUndefined, ErrReasonCodeInvalid
	}

	return r, nil
}

// Validate ensures a reason is valid.
func (r Reason) Validate() error {
	if _, ok := ReasonName[r]; !ok {
		return ErrReasonCodeInvalid
	}

	return nil
}

// Identifier returns the identifier.
func (r Reason) Identifier() string {
	id, ok := ReasonName[r]
	if !ok {
		return fmt.Sprintf("Reason(%d)", r)
	}

	return id
}

// Description returns the description of a reason code, along with an error if
// it is an unknown value.
func (r Reason) Description() (string, error) {
	s, ok := reasonDescriptions[r]
	if !ok {
		return fmt.Sprintf("%%!Reason(%d)", r), fmt.Errorf("reason description '%d': %w", r, ErrReasonCodeInvalid)
	}

	return s, nil
}

// String returns a text representation of the reason.
func (r Reason) String() string {
	desc, err := r.Description()
	if err != nil {
		return desc
	}

	return fmt.Sprintf("\"%s: %s\"", r.Identifier(), desc)
}

// GoString returns a text representation of the reason to satisfy the GoStringer
// interface.
func (r Reason) GoString() string {
	return fmt.Sprintf("%%!Reason(%d)", r)
}

// Deprecated indicates if this method is a deprecated method.
//
//nolint:gocritic
func (r Reason) Deprecated() bool {
	dep, ok := reasonsDeprecated[r]
	if !ok {
		return false
	}

	return dep
}

// RequiresDescription flag the need for custom text for description.
func (r Reason) RequiresDescription() bool {
	return r == ReasonFreeTextDescription
}
