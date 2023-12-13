package clickhouse

import (
	"context"

	"go.uber.org/zap"

	//"github.com/docker/docker/daemon/logger"
	devicepb "github.com/irisco88/protos/gen/device/v1"
	//"go.uber.org/zap"
	"strings"
)

func (adb *AVLDataBase) GetLastPointsData(ctx context.Context, dataFilter string) ([]*devicepb.AVLData, error) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	dataFilterQ := "  "
	hasOR := " and "
	whereQuery := ""
	dataSpeed := ""
	dataIgnition := ""
	dataGSMSignal := ""
	dataDigitalInput1 := ""
	dataDigitalOutput1 := ""
	dataSDStatus := ""
	dataDigitalInput2 := ""
	dataDigitalOutput2 := ""
	dataCrashDetection := ""
	dataOverSpeeding := ""
	dataExternalVoltage := ""
	dataBatteryVoltage := ""
	dataAnalogInput1 := ""
	dataAnalogInput2 := ""
	dataAnalogInput3 := ""
	dataAnalogInput4 := ""
	dataPCBTemperature := ""
	dataVehicleSpeed := ""
	dataEngineSpeed_RPM := ""
	dataEngineCoolantTemperature := ""
	dataFuellevelintank := ""
	dataCheckEngine := ""
	dataAirConditionPressureSwitch1 := ""
	dataAirConditionPressureSwitch2 := ""
	dataGearShiftindicator := ""
	dataDesiredGearValue := ""
	dataVehicleType := ""
	dataConditionimmobilizer := ""
	dataBrakePedalStatus := ""
	dataClutchPedalStatus := ""
	dataGearEngagedStatus := ""
	dataActualAccPedal := ""
	dataEngineThrottlePosition := ""
	dataIndicatedEngineTorque := ""
	dataEngineFrictionTorque := ""
	dataEngineActualTorque := ""
	dataCruiseControlOn_Off := ""
	dataSpeedLimiterOn_Off := ""
	dataConditioncruisecontrollamp := ""
	dataEngineFuleCutOff := ""
	dataConditioncatalystheatingactivated := ""
	dataACcompressorstatus := ""
	dataConditionmainrelay := ""
	datadistance := ""
	dataVirtualAccPedal := ""
	dataIntakeairtemperature := ""
	dataDesiredSpeed := ""
	dataOiltemperature_TCU := ""
	dataAmbientairtemperature := ""
	dataEMS_DTC := ""
	dataABS_DTC := ""
	dataBCM_DTC := ""
	dataACU_DTC := ""
	dataESC_DTC := ""
	dataICN_DTC := ""
	dataEPS_DTC := ""
	dataCAS_DTC := ""
	dataFCM_FN_DTC := ""
	dataICU_DTC := ""
	dataReserve_DTC := ""
	dataSensor1 := ""
	dataSensor2 := ""
	dataSensor3 := ""
	dataSensor4 := ""
	dataSensor5 := ""
	dataSensor6 := ""
	dataSensor7 := ""
	dataSensor8 := ""
	dataSensor9 := ""
	dataSensor10 := ""
	dataSensor11 := ""
	dataSensor12 := ""
	dataSensor13 := ""
	dataSensor14 := ""
	dataSensor15 := ""
	dataSensor16 := ""
	dataSensor17 := ""
	dataSensor18 := ""
	dataSensor19 := ""
	dataSensor20 := ""
	if strings.Contains(dataFilter, "hasOR") {
		hasOR = " or "
		whereQuery = " where imei=='0' "
	} else {
		whereQuery = " where imei!='0' "
	}
	if strings.Contains(dataFilter, "speed=") ||
		strings.Contains(dataFilter, "Ignition=") ||
		strings.Contains(dataFilter, "GSMSignal=") ||
		strings.Contains(dataFilter, "DigitalInput1=") ||
		strings.Contains(dataFilter, "DigitalOutput1=") ||
		strings.Contains(dataFilter, "SDStatus=") ||
		strings.Contains(dataFilter, "DigitalInput2=") ||
		strings.Contains(dataFilter, "DigitalOutput2=") ||
		strings.Contains(dataFilter, "CrashDetection=") ||
		strings.Contains(dataFilter, "OverSpeeding=") ||
		strings.Contains(dataFilter, "ExternalVoltage=") ||
		strings.Contains(dataFilter, "BatteryVoltage=") ||
		strings.Contains(dataFilter, "AnalogInput1=") ||
		strings.Contains(dataFilter, "AnalogInput2=") ||
		strings.Contains(dataFilter, "AnalogInput3=") ||
		strings.Contains(dataFilter, "AnalogInput4=") ||
		strings.Contains(dataFilter, "PCBTemperature=") ||
		strings.Contains(dataFilter, "VehicleSpeed=") ||
		strings.Contains(dataFilter, "EngineSpeed_RPM=") ||
		strings.Contains(dataFilter, "EngineCoolantTemperature=") ||
		strings.Contains(dataFilter, "FuelLevelinTank=") ||
		strings.Contains(dataFilter, "CheckEngine=") ||
		strings.Contains(dataFilter, "AirConditionPressureSwitch1=") ||
		strings.Contains(dataFilter, "AirConditionPressureSwitch2=") ||
		strings.Contains(dataFilter, "GearShiftindicator=") ||
		strings.Contains(dataFilter, "DesiredGearValue=") ||
		strings.Contains(dataFilter, "VehicleType=") ||
		strings.Contains(dataFilter, "Conditionimmobilizer=") ||
		strings.Contains(dataFilter, "BrakePedalStatus=") ||
		strings.Contains(dataFilter, "ClutchPedalStatus=") ||
		strings.Contains(dataFilter, "GearEngagedStatus=") ||
		strings.Contains(dataFilter, "ActualAccPedal=") ||
		strings.Contains(dataFilter, "EngineThrottlePosition=") ||
		strings.Contains(dataFilter, "IndicatedEngineTorque=") ||
		strings.Contains(dataFilter, "EngineFrictionTorque=") ||
		strings.Contains(dataFilter, "EngineActualTorque=") ||
		strings.Contains(dataFilter, "CruiseControlOn_Off=") ||
		strings.Contains(dataFilter, "SpeedLimiterOn_Off=") ||
		strings.Contains(dataFilter, "conditionCruisControlLamp=") ||
		strings.Contains(dataFilter, "EngineFuleCutOff=") ||
		strings.Contains(dataFilter, "ConditionCatalystHeatingActivated=") ||
		strings.Contains(dataFilter, "ACcompressorstatus=") ||
		strings.Contains(dataFilter, "ConditionMainRelay=") ||
		strings.Contains(dataFilter, "distance=") ||
		strings.Contains(dataFilter, "VirtualAccPeda=") ||
		strings.Contains(dataFilter, "Intakeairtemperature=") ||
		strings.Contains(dataFilter, "DesiredSpeed=") ||
		strings.Contains(dataFilter, "OilTemperature(TCU)=") ||
		strings.Contains(dataFilter, "AmbientAirTemperature=") ||
		strings.Contains(dataFilter, "EMS_DTC=") ||
		strings.Contains(dataFilter, "ABS_DTC=") ||
		strings.Contains(dataFilter, "BCM_DTC=") ||
		strings.Contains(dataFilter, "ACU_DTC=") ||
		strings.Contains(dataFilter, "ESC_DTC=") ||
		strings.Contains(dataFilter, "ICN_DTC=") ||
		strings.Contains(dataFilter, "EPS_DTC=") ||
		strings.Contains(dataFilter, "CAS_DTC=") ||
		strings.Contains(dataFilter, "FCM/FN_DTC=") ||
		strings.Contains(dataFilter, "ICU_DTC=") ||
		strings.Contains(dataFilter, "Reserve_DTC=") ||
		strings.Contains(dataFilter, "Sensor1=") ||
		strings.Contains(dataFilter, "Sensor2=") ||
		strings.Contains(dataFilter, "Sensor3=") ||
		strings.Contains(dataFilter, "Sensor4=") ||
		strings.Contains(dataFilter, "Sensor5=") ||
		strings.Contains(dataFilter, "Sensor6=") ||
		strings.Contains(dataFilter, "Sensor7=") ||
		strings.Contains(dataFilter, "Sensor8=") ||
		strings.Contains(dataFilter, "Sensor9=") ||
		strings.Contains(dataFilter, "Sensor10=") ||
		strings.Contains(dataFilter, "Sensor11=") ||
		strings.Contains(dataFilter, "Sensor12=") ||
		strings.Contains(dataFilter, "Sensor13=") ||
		strings.Contains(dataFilter, "Sensor14=") ||
		strings.Contains(dataFilter, "Sensor15=") ||
		strings.Contains(dataFilter, "Sensor16=") ||
		strings.Contains(dataFilter, "Sensor17=") ||
		strings.Contains(dataFilter, "Sensor18=") ||
		strings.Contains(dataFilter, "Sensor19=") ||
		strings.Contains(dataFilter, "Sensor20=") {
		if strings.Contains(dataFilter, "speed=") {
			speeds := strings.Split(dataFilter, "speed=")[1]
			speed := strings.Split(speeds, "&")[0]
			speed1 := strings.Split(speed, "_")[0]
			speed2 := strings.Split(speed, "_")[1]
			dataSpeed = hasOR + "  speed  >=  " + speed1 + hasOR + " speed <= " + speed2
		}
		if strings.Contains(dataFilter, "Ignition=") {
			Ignitions := strings.Split(dataFilter, "Ignition=")[1]
			Ignition := strings.Split(Ignitions, "&")[0]
			Ignition1 := strings.Split(Ignition, "_")[0]
			Ignition2 := strings.Split(Ignition, "_")[1]
			dataIgnition = hasOR + " io_elements['Ignition']  >=  " + Ignition1 + hasOR + "  io_elements['Ignition'] <=  " + Ignition2

		}
		if strings.Contains(dataFilter, "GSMSignal=") {
			GSMSignals := strings.Split(dataFilter, "GSMSignal=")[1]
			GSMSignal := strings.Split(GSMSignals, "&")[0]
			GSMSignal1 := strings.Split(GSMSignal, "_")[0]
			GSMSignal2 := strings.Split(GSMSignal, "_")[1]
			dataGSMSignal = hasOR + "  io_elements['GSMSignal']  >=  " + GSMSignal1 + hasOR + "  io_elements['GSMSignal'] <= " + GSMSignal2
		}
		if strings.Contains(dataFilter, "DigitalInput1=") {
			DigitalInput1s := strings.Split(dataFilter, "DigitalInput1=")[1]
			DigitalInput1 := strings.Split(DigitalInput1s, "&")[0]
			DigitalInput11 := strings.Split(DigitalInput1, "_")[0]
			DigitalInput12 := strings.Split(DigitalInput1, "_")[1]
			dataDigitalInput1 = hasOR + "  io_elements['DigitalInput1']  >=  " + DigitalInput11 + hasOR + " io_elements['DigitalInput1'] <= " + DigitalInput12
		}
		if strings.Contains(dataFilter, "DigitalOutput1=") {
			DigitalOutput1s := strings.Split(dataFilter, "DigitalOutput1=")[1]
			DigitalOutput1 := strings.Split(DigitalOutput1s, "&")[0]
			DigitalOutput11 := strings.Split(DigitalOutput1, "_")[0]
			DigitalOutput12 := strings.Split(DigitalOutput1, "_")[1]
			dataDigitalOutput1 = hasOR + " io_elements['DigitalOutput1']  >=  " + DigitalOutput11 + hasOR + " io_elements['DigitalOutput1'] <= " + DigitalOutput12
		}
		if strings.Contains(dataFilter, "SDStatus=") {
			SDStatuss := strings.Split(dataFilter, "SDStatus=")[1]
			SDStatus := strings.Split(SDStatuss, "&")[0]
			SDStatus1 := strings.Split(SDStatus, "_")[0]
			SDStatus2 := strings.Split(SDStatus, "_")[1]
			dataSDStatus = hasOR + " io_elements['SDStatus']  >=  " + SDStatus1 + hasOR + "  io_elements['SDStatus'] <= " + SDStatus2
		}
		if strings.Contains(dataFilter, "DigitalInput2=") {
			DigitalInput2s := strings.Split(dataFilter, "DigitalInput2=")[1]
			DigitalInput2 := strings.Split(DigitalInput2s, "&")[0]
			DigitalInput21 := strings.Split(DigitalInput2, "_")[0]
			DigitalInput22 := strings.Split(DigitalInput2, "_")[1]
			dataDigitalInput2 = hasOR + " io_elements['DigitalInput2']  >=  " + DigitalInput21 + hasOR + "  io_elements['DigitalInput2'] <= " + DigitalInput22
		}
		if strings.Contains(dataFilter, "DigitalOutput2=") {
			DigitalOutput2s := strings.Split(dataFilter, "DigitalOutput2=")[1]
			DigitalOutput2 := strings.Split(DigitalOutput2s, "&")[0]
			DigitalOutput21 := strings.Split(DigitalOutput2, "_")[0]
			DigitalOutput22 := strings.Split(DigitalOutput2, "_")[1]
			dataDigitalOutput2 = hasOR + "  io_elements['DigitalOutput2']  >=  " + DigitalOutput21 + hasOR + "  io_elements['DigitalOutput2'] <=  " + DigitalOutput22
		}
		if strings.Contains(dataFilter, "CrashDetection=") {
			CrashDetections := strings.Split(dataFilter, "CrashDetection=")[1]
			CrashDetection := strings.Split(CrashDetections, "&")[0]
			CrashDetection1 := strings.Split(CrashDetection, "_")[0]
			CrashDetection2 := strings.Split(CrashDetection, "_")[1]
			dataCrashDetection = hasOR + "  io_elements['CrashDetection']  >=  " + CrashDetection1 + hasOR + "  io_elements['CrashDetection'] <= " + CrashDetection2
		}
		if strings.Contains(dataFilter, "OverSpeeding=") {
			OverSpeedings := strings.Split(dataFilter, "OverSpeeding=")[1]
			OverSpeeding := strings.Split(OverSpeedings, "&")[0]
			OverSpeeding1 := strings.Split(OverSpeeding, "_")[0]
			OverSpeeding2 := strings.Split(OverSpeeding, "_")[1]
			dataOverSpeeding = hasOR + " io_elements['OverSpeeding']  >=  " + OverSpeeding1 + hasOR + "   io_elements['OverSpeeding'] <= " + OverSpeeding2
		}
		if strings.Contains(dataFilter, "ExternalVoltage=") {
			ExternalVoltages := strings.Split(dataFilter, "ExternalVoltage=")[1]
			ExternalVoltage := strings.Split(ExternalVoltages, "&")[0]
			ExternalVoltage1 := strings.Split(ExternalVoltage, "_")[0]
			ExternalVoltage2 := strings.Split(ExternalVoltage, "_")[1]
			dataExternalVoltage = hasOR + " io_elements['ExternalVoltage']  >=  " + ExternalVoltage1 + hasOR + "  io_elements['ExternalVoltage'] <= " + ExternalVoltage2
		}
		if strings.Contains(dataFilter, "BatteryVoltage=") {
			BatteryVoltages := strings.Split(dataFilter, "BatteryVoltage=")[1]
			BatteryVoltage := strings.Split(BatteryVoltages, "&")[0]
			BatteryVoltage1 := strings.Split(BatteryVoltage, "_")[0]
			BatteryVoltage2 := strings.Split(BatteryVoltage, "_")[1]
			dataBatteryVoltage = hasOR + " io_elements['BatteryVoltage']  >=  " + BatteryVoltage1 + hasOR + "  io_elements['BatteryVoltage'] <= " + BatteryVoltage2
		}
		if strings.Contains(dataFilter, "AnalogInput1=") {
			AnalogInput1s := strings.Split(dataFilter, "AnalogInput1=")[1]
			AnalogInput1 := strings.Split(AnalogInput1s, "&")[0]
			AnalogInput11 := strings.Split(AnalogInput1, "_")[0]
			AnalogInput12 := strings.Split(AnalogInput1, "_")[1]
			dataAnalogInput1 = hasOR + " io_elements['AnalogInput1']  >=  " + AnalogInput11 + hasOR + " io_elements['AnalogInput1'] <= " + AnalogInput12
		}
		if strings.Contains(dataFilter, "AnalogInput2=") {
			AnalogInput2s := strings.Split(dataFilter, "AnalogInput2=")[1]
			AnalogInput2 := strings.Split(AnalogInput2s, "&")[0]
			AnalogInput21 := strings.Split(AnalogInput2, "_")[0]
			AnalogInput22 := strings.Split(AnalogInput2, "_")[1]
			dataAnalogInput2 = hasOR + "  io_elements['AnalogInput2']  >=  " + AnalogInput21 + hasOR + "  io_elements['AnalogInput2'] <= " + AnalogInput22
		}
		if strings.Contains(dataFilter, "AnalogInput3=") {
			AnalogInput3s := strings.Split(dataFilter, "AnalogInput3=")[1]
			AnalogInput3 := strings.Split(AnalogInput3s, "&")[0]
			AnalogInput31 := strings.Split(AnalogInput3, "_")[0]
			AnalogInput32 := strings.Split(AnalogInput3, "_")[1]
			dataAnalogInput3 = hasOR + " io_elements['AnalogInput3']  >=  " + AnalogInput31 + hasOR + "  io_elements['AnalogInput3'] <= " + AnalogInput32
		}
		if strings.Contains(dataFilter, "AnalogInput4=") {
			AnalogInput4s := strings.Split(dataFilter, "AnalogInput4=")[1]
			AnalogInput4 := strings.Split(AnalogInput4s, "&")[0]
			AnalogInput41 := strings.Split(AnalogInput4, "_")[0]
			AnalogInput42 := strings.Split(AnalogInput4, "_")[1]
			dataAnalogInput4 = hasOR + " io_elements['AnalogInput4']  >=  " + AnalogInput41 + hasOR + "   io_elements['AnalogInput4'] <= " + AnalogInput42
		}
		if strings.Contains(dataFilter, "PCBTemperature=") {
			PCBTemperatures := strings.Split(dataFilter, "PCBTemperature=")[1]
			PCBTemperature := strings.Split(PCBTemperatures, "&")[0]
			PCBTemperature1 := strings.Split(PCBTemperature, "_")[0]
			PCBTemperature2 := strings.Split(PCBTemperature, "_")[1]
			dataPCBTemperature = hasOR + " io_elements['PCBTemperature']  >=  " + PCBTemperature1 + hasOR + "  io_elements['PCBTemperature'] <= " + PCBTemperature2
		}
		if strings.Contains(dataFilter, "VehicleSpeed=") {
			VehicleSpeeds := strings.Split(dataFilter, "VehicleSpeed=")[1]
			VehicleSpeed := strings.Split(VehicleSpeeds, "&")[0]
			VehicleSpeed1 := strings.Split(VehicleSpeed, "_")[0]
			VehicleSpeed2 := strings.Split(VehicleSpeed, "_")[1]
			dataVehicleSpeed = hasOR + " io_elements['VehicleSpeed']  >=  " + VehicleSpeed1 + hasOR + "  io_elements['VehicleSpeed'] <= " + VehicleSpeed2
		}
		if strings.Contains(dataFilter, "EngineSpeed_RPM=") {
			EngineSpeed_RPMs := strings.Split(dataFilter, "EngineSpeed_RPM=")[1]
			EngineSpeed_RPM := strings.Split(EngineSpeed_RPMs, "&")[0]
			EngineSpeed_RPM1 := strings.Split(EngineSpeed_RPM, "_")[0]
			EngineSpeed_RPM2 := strings.Split(EngineSpeed_RPM, "_")[1]
			dataEngineSpeed_RPM = hasOR + " io_elements['EngineSpeed_RPM']  >=  " + EngineSpeed_RPM1 + hasOR + "  io_elements['EngineSpeed_RPM'] <= " + EngineSpeed_RPM2
		}
		if strings.Contains(dataFilter, "EngineCoolantTemperature=") {
			EngineCoolantTemperatures := strings.Split(dataFilter, "EngineCoolantTemperature=")[1]
			EngineCoolantTemperature := strings.Split(EngineCoolantTemperatures, "&")[0]
			EngineCoolantTemperature1 := strings.Split(EngineCoolantTemperature, "_")[0]
			EngineCoolantTemperature2 := strings.Split(EngineCoolantTemperature, "_")[1]
			dataEngineCoolantTemperature = hasOR + " io_elements['EngineCoolantTemperature']  >=  " + EngineCoolantTemperature1 + hasOR + " io_elements['EngineCoolantTemperature'] <= " + EngineCoolantTemperature2
		}
		if strings.Contains(dataFilter, "FuelLevelinTank=") {
			Fuellevelintanks := strings.Split(dataFilter, "FuelLevelinTank=")[1]
			Fuellevelintank := strings.Split(Fuellevelintanks, "&")[0]
			Fuellevelintank1 := strings.Split(Fuellevelintank, "_")[0]
			Fuellevelintank2 := strings.Split(Fuellevelintank, "_")[1]
			dataFuellevelintank = hasOR + " io_elements['FuelLevelinTank']  >=  " + Fuellevelintank1 + hasOR + "  io_elements['FuelLevelinTank'] <= " + Fuellevelintank2
		}
		if strings.Contains(dataFilter, "CheckEngine=") {
			CheckEngines := strings.Split(dataFilter, "CheckEngine=")[1]
			CheckEngine := strings.Split(CheckEngines, "&")[0]
			CheckEngine1 := strings.Split(CheckEngine, "_")[0]
			CheckEngine2 := strings.Split(CheckEngine, "_")[1]
			dataCheckEngine = hasOR + "  io_elements['CheckEngine']  >=  " + CheckEngine1 + hasOR + "  io_elements['CheckEngine'] <= " + CheckEngine2
		}
		if strings.Contains(dataFilter, "AirConditionPressureSwitch1=") {
			AirConditionPressureSwitch1s := strings.Split(dataFilter, "AirConditionPressureSwitch1=")[1]
			AirConditionPressureSwitch1 := strings.Split(AirConditionPressureSwitch1s, "&")[0]
			AirConditionPressureSwitch11 := strings.Split(AirConditionPressureSwitch1, "_")[0]
			AirConditionPressureSwitch12 := strings.Split(AirConditionPressureSwitch1, "_")[1]
			dataAirConditionPressureSwitch1 = hasOR + " io_elements['AirConditionPressureSwitch1']  >=  " + AirConditionPressureSwitch11 + hasOR + " io_elements['AirConditionPressureSwitch1'] <= " + AirConditionPressureSwitch12
		}
		if strings.Contains(dataFilter, "AirConditionPressureSwitch2=") {
			AirConditionPressureSwitch2s := strings.Split(dataFilter, "AirConditionPressureSwitch2=")[1]
			AirConditionPressureSwitch2 := strings.Split(AirConditionPressureSwitch2s, "&")[0]
			AirConditionPressureSwitch21 := strings.Split(AirConditionPressureSwitch2, "_")[0]
			AirConditionPressureSwitch22 := strings.Split(AirConditionPressureSwitch2, "_")[1]
			dataAirConditionPressureSwitch2 = hasOR + " io_elements['AirConditionPressureSwitch2']  >=  " + AirConditionPressureSwitch21 + hasOR + " io_elements['AirConditionPressureSwitch2'] <= " + AirConditionPressureSwitch22
		}
		if strings.Contains(dataFilter, "GearShiftindicator=") {
			GearShiftindicators := strings.Split(dataFilter, "GearShiftindicator=")[1]
			GearShiftindicator := strings.Split(GearShiftindicators, "&")[0]
			GearShiftindicator1 := strings.Split(GearShiftindicator, "_")[0]
			GearShiftindicator2 := strings.Split(GearShiftindicator, "_")[1]
			dataGearShiftindicator = hasOR + " io_elements['GearShiftindicator']  >=  " + GearShiftindicator1 + hasOR + "  io_elements['GearShiftindicator'] <= " + GearShiftindicator2
		}
		if strings.Contains(dataFilter, "DesiredGearValue=") {
			DesiredGearValues := strings.Split(dataFilter, "DesiredGearValue=")[1]
			DesiredGearValue := strings.Split(DesiredGearValues, "&")[0]
			DesiredGearValue1 := strings.Split(DesiredGearValue, "_")[0]
			DesiredGearValue2 := strings.Split(DesiredGearValue, "_")[1]
			dataDesiredGearValue = hasOR + " io_elements['DesiredGearValue']  >=  " + DesiredGearValue1 + hasOR + "  io_elements['DesiredGearValue'] <= " + DesiredGearValue2
		}
		if strings.Contains(dataFilter, "VehicleType=") {
			VehicleTypes := strings.Split(dataFilter, "VehicleType=")[1]
			VehicleType := strings.Split(VehicleTypes, "&")[0]
			VehicleType1 := strings.Split(VehicleType, "_")[0]
			VehicleType2 := strings.Split(VehicleType, "_")[1]
			dataVehicleType = hasOR + " io_elements['VehicleType']  >=  " + VehicleType1 + hasOR + "  io_elements['VehicleType'] <= " + VehicleType2

		}
		if strings.Contains(dataFilter, "Conditionimmobilizer=") {
			Conditionimmobilizers := strings.Split(dataFilter, "Conditionimmobilizer=")[1]
			Conditionimmobilizer := strings.Split(Conditionimmobilizers, "&")[0]
			Conditionimmobilizer1 := strings.Split(Conditionimmobilizer, "_")[0]
			Conditionimmobilizer2 := strings.Split(Conditionimmobilizer, "_")[1]
			dataConditionimmobilizer = hasOR + "  io_elements['Conditionimmobilizer']  >=  " + Conditionimmobilizer1 + hasOR + "   io_elements['Conditionimmobilizer'] <= " + Conditionimmobilizer2
		}
		if strings.Contains(dataFilter, "BrakePedalStatus=") {
			BrakePedalStatuss := strings.Split(dataFilter, "BrakePedalStatus=")[1]
			BrakePedalStatus := strings.Split(BrakePedalStatuss, "&")[0]
			BrakePedalStatus1 := strings.Split(BrakePedalStatus, "_")[0]
			BrakePedalStatus2 := strings.Split(BrakePedalStatus, "_")[1]
			dataBrakePedalStatus = hasOR + " io_elements['BrakePedalStatus']  >=  " + BrakePedalStatus1 + hasOR + "  io_elements['BrakePedalStatus'] <= " + BrakePedalStatus2
		}
		if strings.Contains(dataFilter, "ClutchPedalStatus=") {
			ClutchPedalStatuss := strings.Split(dataFilter, "ClutchPedalStatus=")[1]
			ClutchPedalStatus := strings.Split(ClutchPedalStatuss, "&")[0]
			ClutchPedalStatus1 := strings.Split(ClutchPedalStatus, "_")[0]
			ClutchPedalStatus2 := strings.Split(ClutchPedalStatus, "_")[1]
			dataClutchPedalStatus = hasOR + " io_elements['ClutchPedalStatus']  >=  " + ClutchPedalStatus1 + hasOR + "   io_elements['ClutchPedalStatus'] <= " + ClutchPedalStatus2
		}
		if strings.Contains(dataFilter, "GearEngagedStatus=") {
			GearEngagedStatuss := strings.Split(dataFilter, "GearEngagedStatus=")[1]
			GearEngagedStatus := strings.Split(GearEngagedStatuss, "&")[0]
			GearEngagedStatus1 := strings.Split(GearEngagedStatus, "_")[0]
			GearEngagedStatus2 := strings.Split(GearEngagedStatus, "_")[1]
			dataGearEngagedStatus = hasOR + " io_elements['GearEngagedStatus']  >=  " + GearEngagedStatus1 + hasOR + "  io_elements['GearEngagedStatus'] <= " + GearEngagedStatus2
		}
		if strings.Contains(dataFilter, "ActualAccPedal=") {
			ActualAccPedals := strings.Split(dataFilter, "ActualAccPedal=")[1]
			ActualAccPedal := strings.Split(ActualAccPedals, "&")[0]
			ActualAccPedal1 := strings.Split(ActualAccPedal, "_")[0]
			ActualAccPedal2 := strings.Split(ActualAccPedal, "_")[1]
			dataActualAccPedal = hasOR + "  io_elements['ActualAccPedal']  >=  " + ActualAccPedal1 + hasOR + "   io_elements['ActualAccPedal'] <= " + ActualAccPedal2
		}
		if strings.Contains(dataFilter, "EngineThrottlePosition=") {
			EngineThrottlePositions := strings.Split(dataFilter, "EngineThrottlePosition=")[1]
			EngineThrottlePosition := strings.Split(EngineThrottlePositions, "&")[0]
			EngineThrottlePosition1 := strings.Split(EngineThrottlePosition, "_")[0]
			EngineThrottlePosition2 := strings.Split(EngineThrottlePosition, "_")[1]
			dataEngineThrottlePosition = hasOR + "  io_elements['EngineThrottlePosition']  >=  " + EngineThrottlePosition1 + hasOR + "   io_elements['EngineThrottlePosition'] <= " + EngineThrottlePosition2
		}
		if strings.Contains(dataFilter, "IndicatedEngineTorque=") {
			IndicatedEngineTorques := strings.Split(dataFilter, "IndicatedEngineTorque=")[1]
			IndicatedEngineTorque := strings.Split(IndicatedEngineTorques, "&")[0]
			IndicatedEngineTorque1 := strings.Split(IndicatedEngineTorque, "_")[0]
			IndicatedEngineTorque2 := strings.Split(IndicatedEngineTorque, "_")[1]
			dataIndicatedEngineTorque = hasOR + "  io_elements['IndicatedEngineTorque']  >=  " + IndicatedEngineTorque1 + hasOR + "   io_elements['IndicatedEngineTorque'] <= " + IndicatedEngineTorque2
		}
		if strings.Contains(dataFilter, "EngineFrictionTorque=") {
			EngineFrictionTorques := strings.Split(dataFilter, "EngineFrictionTorque=")[1]
			EngineFrictionTorque := strings.Split(EngineFrictionTorques, "&")[0]
			EngineFrictionTorque1 := strings.Split(EngineFrictionTorque, "_")[0]
			EngineFrictionTorque2 := strings.Split(EngineFrictionTorque, "_")[1]
			dataEngineFrictionTorque = hasOR + " io_elements['EngineFrictionTorque']  >=  " + EngineFrictionTorque1 + hasOR + "     io_elements['EngineFrictionTorque'] <= " + EngineFrictionTorque2
		}
		if strings.Contains(dataFilter, "EngineActualTorque=") {
			EngineActualTorques := strings.Split(dataFilter, "EngineActualTorque=")[1]
			EngineActualTorque := strings.Split(EngineActualTorques, "&")[0]
			EngineActualTorque1 := strings.Split(EngineActualTorque, "_")[0]
			EngineActualTorque2 := strings.Split(EngineActualTorque, "_")[1]
			dataEngineActualTorque = hasOR + "  io_elements['EngineActualTorque']  >=  " + EngineActualTorque1 + hasOR + "    io_elements['EngineActualTorque'] <= " + EngineActualTorque2
		}
		if strings.Contains(dataFilter, "CruiseControlOn_Off=") {
			CruiseControlOn_Offs := strings.Split(dataFilter, "CruiseControlOn_Off=")[1]
			CruiseControlOn_Off := strings.Split(CruiseControlOn_Offs, "&")[0]
			CruiseControlOn_Off1 := strings.Split(CruiseControlOn_Off, "_")[0]
			CruiseControlOn_Off2 := strings.Split(CruiseControlOn_Off, "_")[1]
			dataCruiseControlOn_Off = hasOR + "  io_elements['CruiseControlOn_Off']  >=  " + CruiseControlOn_Off1 + hasOR + "  io_elements['CruiseControlOn_Off'] <= " + CruiseControlOn_Off2
		}
		if strings.Contains(dataFilter, "SpeedLimiterOn_Off=") {
			SpeedLimiterOn_Offs := strings.Split(dataFilter, "SpeedLimiterOn_Off=")[1]
			SpeedLimiterOn_Off := strings.Split(SpeedLimiterOn_Offs, "&")[0]
			SpeedLimiterOn_Off1 := strings.Split(SpeedLimiterOn_Off, "_")[0]
			SpeedLimiterOn_Off2 := strings.Split(SpeedLimiterOn_Off, "_")[1]
			dataSpeedLimiterOn_Off = hasOR + "  io_elements['SpeedLimiterOn_Off']  >=  " + SpeedLimiterOn_Off1 + hasOR + "   io_elements['SpeedLimiterOn_Off'] <= " + SpeedLimiterOn_Off2
		}
		if strings.Contains(dataFilter, "conditionCruisControlLamp=") {
			conditioncruisecontrollamps := strings.Split(dataFilter, "conditionCruisControlLamp=")[1]
			conditioncruisecontrollamp := strings.Split(conditioncruisecontrollamps, "&")[0]
			conditioncruisecontrollamp1 := strings.Split(conditioncruisecontrollamp, "_")[0]
			conditioncruisecontrollamp2 := strings.Split(conditioncruisecontrollamp, "_")[1]
			dataConditioncruisecontrollamp = hasOR + "  io_elements['conditionCruisControlLamp']  >=  " + conditioncruisecontrollamp1 + hasOR + "    io_elements['conditionCruisControlLamp'] <= " + conditioncruisecontrollamp2
		}
		if strings.Contains(dataFilter, "EngineFuleCutOff=") {
			EngineFuleCutOffs := strings.Split(dataFilter, "EngineFuleCutOff=")[1]
			EngineFuleCutOff := strings.Split(EngineFuleCutOffs, "&")[0]
			EngineFuleCutOff1 := strings.Split(EngineFuleCutOff, "_")[0]
			EngineFuleCutOff2 := strings.Split(EngineFuleCutOff, "_")[1]
			dataEngineFuleCutOff = hasOR + "  io_elements['EngineFuleCutOff']  >=  " + EngineFuleCutOff1 + hasOR + "   io_elements['EngineFuleCutOff'] <= " + EngineFuleCutOff2
		}
		if strings.Contains(dataFilter, "ConditionCatalystHeatingActivated=") {
			Conditioncatalystheatingactivateds := strings.Split(dataFilter, "ConditionCatalystHeatingActivated=")[1]
			Conditioncatalystheatingactivated := strings.Split(Conditioncatalystheatingactivateds, "&")[0]
			Conditioncatalystheatingactivated1 := strings.Split(Conditioncatalystheatingactivated, "_")[0]
			Conditioncatalystheatingactivated2 := strings.Split(Conditioncatalystheatingactivated, "_")[1]
			dataConditioncatalystheatingactivated = hasOR + " io_elements['ConditionCatalystHeatingActivated']  >=  " + Conditioncatalystheatingactivated1 + hasOR + " io_elements['ConditionCatalystHeatingActivated'] <=  " + Conditioncatalystheatingactivated2
		}
		if strings.Contains(dataFilter, "ACcompressorstatus=") {
			ACcompressorstatuss := strings.Split(dataFilter, "ACcompressorstatus=")[1]
			ACcompressorstatus := strings.Split(ACcompressorstatuss, "&")[0]
			ACcompressorstatus1 := strings.Split(ACcompressorstatus, "_")[0]
			ACcompressorstatus2 := strings.Split(ACcompressorstatus, "_")[1]
			dataACcompressorstatus = hasOR + "  io_elements['ACcompressorstatus']  >=  " + ACcompressorstatus1 + hasOR + "   io_elements['ACcompressorstatus'] <= " + ACcompressorstatus2
		}
		if strings.Contains(dataFilter, "ConditionMainRelay=") {
			Conditionmainrelays := strings.Split(dataFilter, "ConditionMainRelay=")[1]
			Conditionmainrelay := strings.Split(Conditionmainrelays, "&")[0]
			Conditionmainrelay1 := strings.Split(Conditionmainrelay, "_")[0]
			Conditionmainrelay2 := strings.Split(Conditionmainrelay, "_")[1]
			dataConditionmainrelay = hasOR + "  io_elements['Conditionmainrelay']  >=  " + Conditionmainrelay1 + hasOR + "  io_elements['Conditionmainrelay'] <= " + Conditionmainrelay2
		}
		if strings.Contains(dataFilter, "distance=") {
			distances := strings.Split(dataFilter, "distance=")[1]
			distance := strings.Split(distances, "&")[0]
			distance1 := strings.Split(distance, "_")[0]
			distance2 := strings.Split(distance, "_")[1]
			datadistance = hasOR + "  io_elements['distance']  >=  " + distance1 + hasOR + "  io_elements['distance'] <= " + distance2
		}
		if strings.Contains(dataFilter, "VirtualAccPeda=") {
			VirtualAccPedals := strings.Split(dataFilter, "VirtualAccPedal=")[1]
			VirtualAccPedal := strings.Split(VirtualAccPedals, "&")[0]
			VirtualAccPedal1 := strings.Split(VirtualAccPedal, "_")[0]
			VirtualAccPedal2 := strings.Split(VirtualAccPedal, "_")[1]
			dataVirtualAccPedal = hasOR + " io_elements['VirtualAccPedal']  >=  " + VirtualAccPedal1 + hasOR + " io_elements['VirtualAccPedal'] <= " + VirtualAccPedal2
		}
		if strings.Contains(dataFilter, "Intakeairtemperature=") {
			Intakeairtemperatures := strings.Split(dataFilter, "Intakeairtemperature=")[1]
			Intakeairtemperature := strings.Split(Intakeairtemperatures, "&")[0]
			Intakeairtemperature1 := strings.Split(Intakeairtemperature, "_")[0]
			Intakeairtemperature2 := strings.Split(Intakeairtemperature, "_")[1]
			dataIntakeairtemperature = hasOR + " io_elements['Intakeairtemperature']  >=  " + Intakeairtemperature1 + hasOR + "  io_elements['Intakeairtemperature'] <= " + Intakeairtemperature2
		}
		if strings.Contains(dataFilter, "DesiredSpeed=") {
			DesiredSpeeds := strings.Split(dataFilter, "DesiredSpeed=")[1]
			DesiredSpeed := strings.Split(DesiredSpeeds, "&")[0]
			DesiredSpeed1 := strings.Split(DesiredSpeed, "_")[0]
			DesiredSpeed2 := strings.Split(DesiredSpeed, "_")[1]
			dataDesiredSpeed = hasOR + "  io_elements['DesiredSpeed']  >=  " + DesiredSpeed1 + hasOR + " io_elements['DesiredSpeed'] <= " + DesiredSpeed2
		}
		if strings.Contains(dataFilter, "OilTemperature(TCU)=") {
			Oiltemperature_TCUs := strings.Split(dataFilter, "OilTemperature(TCU)=")[1]
			Oiltemperature_TCU := strings.Split(Oiltemperature_TCUs, "&")[0]
			Oiltemperature_TCU1 := strings.Split(Oiltemperature_TCU, "_")[0]
			Oiltemperature_TCU2 := strings.Split(Oiltemperature_TCU, "_")[1]
			dataOiltemperature_TCU = hasOR + "  io_elements['OilTemperature(TCU)']  >=  " + Oiltemperature_TCU1 + hasOR + "  io_elements['OilTemperature(TCU)'] <=  " + Oiltemperature_TCU2
		}
		if strings.Contains(dataFilter, "AmbientAirTemperature=") {
			Ambientairtemperatures := strings.Split(dataFilter, "AmbientAirTemperature=")[1]
			Ambientairtemperature := strings.Split(Ambientairtemperatures, "&")[0]
			Ambientairtemperature1 := strings.Split(Ambientairtemperature, "_")[0]
			Ambientairtemperature2 := strings.Split(Ambientairtemperature, "_")[1]
			dataAmbientairtemperature = hasOR + " io_elements['AmbientAirTemperature']  >=  " + Ambientairtemperature1 + hasOR + "  io_elements['AmbientAirTemperature'] <= " + Ambientairtemperature2
		}
		if strings.Contains(dataFilter, "EMS_DTC=") {
			EMS_DTCs := strings.Split(dataFilter, "EMS_DTC=")[1]
			EMS_DTC := strings.Split(EMS_DTCs, "&")[0]
			EMS_DTC1 := strings.Split(EMS_DTC, "_")[0]
			EMS_DTC2 := strings.Split(EMS_DTC, "_")[1]
			dataEMS_DTC = hasOR + "  io_elements['EMS_DTC']  >=  " + EMS_DTC1 + hasOR + " io_elements['EMS_DTC'] <= " + EMS_DTC2
		}
		if strings.Contains(dataFilter, "ABS_DTC=") {
			ABS_DTCs := strings.Split(dataFilter, "ABS_DTC=")[1]
			ABS_DTC := strings.Split(ABS_DTCs, "&")[0]
			ABS_DTC1 := strings.Split(ABS_DTC, "_")[0]
			ABS_DTC2 := strings.Split(ABS_DTC, "_")[1]
			dataABS_DTC = hasOR + "  io_elements['ABS_DTC']  >=  " + ABS_DTC1 + hasOR + "   io_elements['ABS_DTC'] <= " + ABS_DTC2
		}
		if strings.Contains(dataFilter, "BCM_DTC=") {
			BCM_DTCs := strings.Split(dataFilter, "BCM_DTC=")[1]
			BCM_DTC := strings.Split(BCM_DTCs, "&")[0]
			BCM_DTC1 := strings.Split(BCM_DTC, "_")[0]
			BCM_DTC2 := strings.Split(BCM_DTC, "_")[1]
			dataBCM_DTC = hasOR + "  io_elements['BCM_DTC']  >=  " + BCM_DTC1 + hasOR + "  io_elements['BCM_DTC'] <= " + BCM_DTC2
		}
		if strings.Contains(dataFilter, "ACU_DTC=") {
			ACU_DTCs := strings.Split(dataFilter, "ACU_DTC=")[1]
			ACU_DTC := strings.Split(ACU_DTCs, "&")[0]
			ACU_DTC1 := strings.Split(ACU_DTC, "_")[0]
			ACU_DTC2 := strings.Split(ACU_DTC, "_")[1]
			dataACU_DTC = hasOR + "  io_elements['ACU_DTC']  >=  " + ACU_DTC1 + hasOR + "    io_elements['ACU_DTC'] <= " + ACU_DTC2
		}
		if strings.Contains(dataFilter, "ESC_DTC=") {
			ESC_DTCs := strings.Split(dataFilter, "ESC_DTC=")[1]
			ESC_DTC := strings.Split(ESC_DTCs, "&")[0]
			ESC_DTC1 := strings.Split(ESC_DTC, "_")[0]
			ESC_DTC2 := strings.Split(ESC_DTC, "_")[1]
			dataESC_DTC = hasOR + "  io_elements['ESC_DTC']  >=  " + ESC_DTC1 + hasOR + " io_elements['ESC_DTC'] <= " + ESC_DTC2
		}
		if strings.Contains(dataFilter, "ICN_DTC=") {
			ICN_DTCs := strings.Split(dataFilter, "ICN_DTC=")[1]
			ICN_DTC := strings.Split(ICN_DTCs, "&")[0]
			ICN_DTC1 := strings.Split(ICN_DTC, "_")[0]
			ICN_DTC2 := strings.Split(ICN_DTC, "_")[1]
			dataICN_DTC = hasOR + "  io_elements['ICN_DTC']  >=  " + ICN_DTC1 + hasOR + "    io_elements['ICN_DTC'] <= " + ICN_DTC2
		}
		if strings.Contains(dataFilter, "EPS_DTC=") {
			EPS_DTCs := strings.Split(dataFilter, "EPS_DTC=")[1]
			EPS_DTC := strings.Split(EPS_DTCs, "&")[0]
			EPS_DTC1 := strings.Split(EPS_DTC, "_")[0]
			EPS_DTC2 := strings.Split(EPS_DTC, "_")[1]
			dataEPS_DTC = hasOR + " io_elements['EPS_DTC']  >=  " + EPS_DTC1 + hasOR + "  io_elements['EPS_DTC'] <= " + EPS_DTC2
		}
		if strings.Contains(dataFilter, "CAS_DTC=") {
			CAS_DTCs := strings.Split(dataFilter, "CAS_DTC=")[1]
			CAS_DTC := strings.Split(CAS_DTCs, "&")[0]
			CAS_DTC1 := strings.Split(CAS_DTC, "_")[0]
			CAS_DTC2 := strings.Split(CAS_DTC, "_")[1]
			dataCAS_DTC = hasOR + " io_elements['CAS_DTC']  >=  " + CAS_DTC1 + hasOR + "  io_elements['CAS_DTC'] <= " + CAS_DTC2
		}
		if strings.Contains(dataFilter, "FCM/FN_DTC=") {
			FCM_FN_DTCs := strings.Split(dataFilter, "FCM/FN_DTC=")[1]
			FCM_FN_DTC := strings.Split(FCM_FN_DTCs, "&")[0]
			FCM_FN_DTC1 := strings.Split(FCM_FN_DTC, "_")[0]
			FCM_FN_DTC2 := strings.Split(FCM_FN_DTC, "_")[1]
			dataFCM_FN_DTC = hasOR + "  io_elements['FCM/FN_DTC']  >=  " + FCM_FN_DTC1 + hasOR + "  io_elements['FCM/FN_DTC'] <= " + FCM_FN_DTC2
		}
		if strings.Contains(dataFilter, "ICU_DTC=") {
			ICU_DTCs := strings.Split(dataFilter, "ICU_DTC=")[1]
			ICU_DTC := strings.Split(ICU_DTCs, "&")[0]
			ICU_DTC1 := strings.Split(ICU_DTC, "_")[0]
			ICU_DTC2 := strings.Split(ICU_DTC, "_")[1]
			dataICU_DTC = hasOR + " io_elements['ICU_DTC']  >=  " + ICU_DTC1 + hasOR + " io_elements['ICU_DTC'] <=  " + ICU_DTC2
		}
		if strings.Contains(dataFilter, "Reserve_DTC=") {
			Reserve_DTCs := strings.Split(dataFilter, "Reserve_DTC=")[1]
			Reserve_DTC := strings.Split(Reserve_DTCs, "&")[0]
			Reserve_DTC1 := strings.Split(Reserve_DTC, "_")[0]
			Reserve_DTC2 := strings.Split(Reserve_DTC, "_")[1]
			dataReserve_DTC = hasOR + " io_elements['Reserve_DTC']  >=  " + Reserve_DTC1 + hasOR + " io_elements['Reserve_DTC'] <= " + Reserve_DTC2
		}
		if strings.Contains(dataFilter, "Sensor1=") {
			Sensor1s := strings.Split(dataFilter, "Sensor1=")[1]
			Sensor1 := strings.Split(Sensor1s, "&")[0]
			Sensor11 := strings.Split(Sensor1, "_")[0]
			Sensor12 := strings.Split(Sensor1, "_")[1]
			dataSensor1 = hasOR + " io_elements['Sensor1']  >=  " + Sensor11 + hasOR + "   io_elements['Sensor1'] <= " + Sensor12
		}
		if strings.Contains(dataFilter, "Sensor2=") {
			Sensor2s := strings.Split(dataFilter, "Sensor2=")[1]
			Sensor2 := strings.Split(Sensor2s, "&")[0]
			Sensor21 := strings.Split(Sensor2, "_")[0]
			Sensor22 := strings.Split(Sensor2, "_")[1]
			dataSensor2 = hasOR + " io_elements['Sensor2']  >=  " + Sensor21 + hasOR + " io_elements['Sensor2'] <= " + Sensor22
		}
		if strings.Contains(dataFilter, "Sensor3=") {
			Sensor3s := strings.Split(dataFilter, "Sensor3=")[1]
			Sensor3 := strings.Split(Sensor3s, "&")[0]
			Sensor31 := strings.Split(Sensor3, "_")[0]
			Sensor32 := strings.Split(Sensor3, "_")[1]
			dataSensor3 = hasOR + "  io_elements['Sensor3']  >=  " + Sensor31 + hasOR + "  io_elements['Sensor3'] <= " + Sensor32
		}
		if strings.Contains(dataFilter, "Sensor4=") {
			Sensor4s := strings.Split(dataFilter, "Sensor4=")[1]
			Sensor4 := strings.Split(Sensor4s, "&")[0]
			Sensor41 := strings.Split(Sensor4, "_")[0]
			Sensor42 := strings.Split(Sensor4, "_")[1]
			dataSensor4 = hasOR + "  io_elements['Sensor4']  >=  " + Sensor41 + hasOR + "  io_elements['Sensor4'] <=  " + Sensor42
		}
		if strings.Contains(dataFilter, "Sensor5=") {
			Sensor5s := strings.Split(dataFilter, "Sensor5=")[1]
			Sensor5 := strings.Split(Sensor5s, "&")[0]
			Sensor51 := strings.Split(Sensor5, "_")[0]
			Sensor52 := strings.Split(Sensor5, "_")[1]
			dataSensor5 = hasOR + "  io_elements['Sensor5']  >=  " + Sensor51 + hasOR + " io_elements['Sensor5'] <=  " + Sensor52
		}
		if strings.Contains(dataFilter, "Sensor6=") {
			Sensor6s := strings.Split(dataFilter, "Sensor6=")[1]
			Sensor6 := strings.Split(Sensor6s, "&")[0]
			Sensor61 := strings.Split(Sensor6, "_")[0]
			Sensor62 := strings.Split(Sensor6, "_")[1]
			dataSensor6 = hasOR + "  io_elements['Sensor6']  >=  " + Sensor61 + hasOR + " io_elements['Sensor6'] <=  " + Sensor62
		}
		if strings.Contains(dataFilter, "Sensor7=") {
			Sensor7s := strings.Split(dataFilter, "Sensor7=")[1]
			Sensor7 := strings.Split(Sensor7s, "&")[0]
			Sensor71 := strings.Split(Sensor7, "_")[0]
			Sensor72 := strings.Split(Sensor7, "_")[1]
			dataSensor7 = hasOR + "  io_elements['Sensor7']  >=  " + Sensor71 + hasOR + "  io_elements['Sensor7'] <= " + Sensor72
		}
		if strings.Contains(dataFilter, "Sensor8=") {
			Sensor8s := strings.Split(dataFilter, "Sensor8=")[1]
			Sensor8 := strings.Split(Sensor8s, "&")[0]
			Sensor81 := strings.Split(Sensor8, "_")[0]
			Sensor82 := strings.Split(Sensor8, "_")[1]
			dataSensor8 = hasOR + "  io_elements['Sensor8']  >=  " + Sensor81 + hasOR + " io_elements['Sensor8'] <= " + Sensor82
		}
		if strings.Contains(dataFilter, "Sensor9=") {
			Sensor9s := strings.Split(dataFilter, "Sensor9=")[1]
			Sensor9 := strings.Split(Sensor9s, "&")[0]
			Sensor91 := strings.Split(Sensor9, "_")[0]
			Sensor92 := strings.Split(Sensor9, "_")[1]
			dataSensor9 = hasOR + "  io_elements['Sensor9']  >=  " + Sensor91 + hasOR + " io_elements['Sensor9'] <= " + Sensor92
		}
		if strings.Contains(dataFilter, "Sensor10=") {
			Sensor10s := strings.Split(dataFilter, "Sensor10=")[1]
			Sensor10 := strings.Split(Sensor10s, "&")[0]
			Sensor101 := strings.Split(Sensor10, "_")[0]
			Sensor102 := strings.Split(Sensor10, "_")[1]
			dataSensor10 = hasOR + "  io_elements['Sensor10']  >=  " + Sensor101 + hasOR + " io_elements['Sensor10'] <= " + Sensor102
		}
		if strings.Contains(dataFilter, "Sensor11=") {
			Sensor11s := strings.Split(dataFilter, "Sensor11=")[1]
			Sensor11 := strings.Split(Sensor11s, "&")[0]
			Sensor111 := strings.Split(Sensor11, "_")[0]
			Sensor112 := strings.Split(Sensor11, "_")[1]
			dataSensor11 = hasOR + "  io_elements['Sensor11']  >=  " + Sensor111 + hasOR + " io_elements['Sensor11'] <= " + Sensor112
		}
		if strings.Contains(dataFilter, "Sensor12=") {
			Sensor12s := strings.Split(dataFilter, "Sensor12=")[1]
			Sensor12 := strings.Split(Sensor12s, "&")[0]
			Sensor121 := strings.Split(Sensor12, "_")[0]
			Sensor122 := strings.Split(Sensor12, "_")[1]
			dataSensor12 = hasOR + "  io_elements['Sensor12']  >=  " + Sensor121 + hasOR + " io_elements['Sensor12'] <=  " + Sensor122
		}
		if strings.Contains(dataFilter, "Sensor13=") {
			Sensor13s := strings.Split(dataFilter, "Sensor13=")[1]
			Sensor13 := strings.Split(Sensor13s, "&")[0]
			Sensor131 := strings.Split(Sensor13, "_")[0]
			Sensor132 := strings.Split(Sensor13, "_")[1]
			dataSensor13 = hasOR + "  io_elements['Sensor13']  >=  " + Sensor131 + hasOR + "  io_elements['Sensor13'] <=  " + Sensor132
		}
		if strings.Contains(dataFilter, "Sensor14=") {
			Sensor14s := strings.Split(dataFilter, "Sensor14=")[1]
			Sensor14 := strings.Split(Sensor14s, "&")[0]
			Sensor141 := strings.Split(Sensor14, "_")[0]
			Sensor142 := strings.Split(Sensor14, "_")[1]
			dataSensor14 = hasOR + "  io_elements['Sensor14']  >=  " + Sensor141 + hasOR + "  io_elements['Sensor14'] <= " + Sensor142
		}
		if strings.Contains(dataFilter, "Sensor15=") {
			Sensor15s := strings.Split(dataFilter, "Sensor15=")[1]
			Sensor15 := strings.Split(Sensor15s, "&")[0]
			Sensor151 := strings.Split(Sensor15, "_")[0]
			Sensor152 := strings.Split(Sensor15, "_")[1]
			dataSensor15 = hasOR + "  io_elements['Sensor15']  >=  " + Sensor151 + hasOR + "   io_elements['Sensor15'] <=  " + Sensor152
		}
		if strings.Contains(dataFilter, "Sensor16=") {
			Sensor16s := strings.Split(dataFilter, "Sensor16=")[1]
			Sensor16 := strings.Split(Sensor16s, "&")[0]
			Sensor161 := strings.Split(Sensor16, "_")[0]
			Sensor162 := strings.Split(Sensor16, "_")[1]
			dataSensor16 = hasOR + "  io_elements['Sensor16']  >=  " + Sensor161 + hasOR + "  io_elements['Sensor16'] <= " + Sensor162
		}
		if strings.Contains(dataFilter, "Sensor17=") {
			Sensor17s := strings.Split(dataFilter, "Sensor17=")[1]
			Sensor17 := strings.Split(Sensor17s, "&")[0]
			Sensor171 := strings.Split(Sensor17, "_")[0]
			Sensor172 := strings.Split(Sensor17, "_")[1]
			dataSensor17 = hasOR + "  io_elements['Sensor17']  >=  " + Sensor171 + hasOR + "  io_elements['Sensor17'] <=  " + Sensor172
		}
		if strings.Contains(dataFilter, "Sensor18=") {
			Sensor18s := strings.Split(dataFilter, "Sensor18=")[1]
			Sensor18 := strings.Split(Sensor18s, "&")[0]
			Sensor181 := strings.Split(Sensor18, "_")[0]
			Sensor182 := strings.Split(Sensor18, "_")[1]
			dataSensor18 = hasOR + " io_elements['Sensor18']  >=  " + Sensor181 + hasOR + " io_elements['Sensor18'] <= " + Sensor182
		}
		if strings.Contains(dataFilter, "Sensor19=") {
			Sensor19s := strings.Split(dataFilter, "Sensor19=")[1]
			Sensor19 := strings.Split(Sensor19s, "&")[0]
			Sensor191 := strings.Split(Sensor19, "_")[0]
			Sensor192 := strings.Split(Sensor19, "_")[1]
			dataSensor19 = hasOR + " io_elements['Sensor19']  >=  " + Sensor191 + hasOR + "  io_elements['Sensor19'] <= " + Sensor192
		}
		if strings.Contains(dataFilter, "Sensor20=") {
			Sensor20s := strings.Split(dataFilter, "Sensor20=")[1]
			Sensor20 := strings.Split(Sensor20s, "&")[0]
			Sensor201 := strings.Split(Sensor20, "_")[0]
			Sensor202 := strings.Split(Sensor20, "_")[1]
			dataSensor20 = hasOR + "  io_elements['Sensor20'] >= " + Sensor201 + hasOR + "  io_elements['Sensor20'] <= " + Sensor202
		}

	}

	if strings.Contains(dataFilter, "from=") || strings.Contains(dataFilter, "to=") || strings.Contains(dataFilter, "imei=") || strings.Contains(dataFilter, "priority=") || strings.Contains(dataFilter, "eventid=") || strings.Contains(dataFilter, "pg=") || strings.Contains(dataFilter, "limit") {
		dataFrom := ""
		dataTo := ""
		dataImei := ""
		dataPriority := ""
		dataEventId := ""
		limitData := ""
		//		dataPg := "1"

		if strings.Contains(dataFilter, "from=") {
			froms := strings.Split(dataFilter, "from=")[1]
			from := strings.Split(froms, "&")[0]
			dataFrom = hasOR + "  timestamp >= '" + from + "' "
		}
		if strings.Contains(dataFilter, "to=") {
			tos := strings.Split(dataFilter, "to=")[1]
			to := strings.Split(tos, "&")[0]
			dataTo = hasOR + " timestamp <= '" + to + "' "
		}
		if strings.Contains(dataFilter, "imei=") {
			imeis := strings.Split(dataFilter, "imei=")[1]
			imei := strings.Split(imeis, "&")[0]
			dataImei = hasOR + " imei in ( " + imei + " ) "
		}
		if strings.Contains(dataFilter, "priority=") {
			prioritys := strings.Split(dataFilter, "priority=")[1]
			priority := strings.Split(prioritys, "&")[0]
			dataPriority = hasOR + " priority in ( " + priority + " ) "
		}
		if strings.Contains(dataFilter, "eventid=") {
			eventids := strings.Split(dataFilter, "eventid=")[1]
			eventid := strings.Split(eventids, "&")[0]
			eventid1 := strings.Split(eventid, "_")[0]
			eventid2 := strings.Split(eventid, "_")[1]
			dataEventId = hasOR + "  event_id >= " + eventid1 + hasOR + "   event_id<= " + eventid2

		}
		//if strings.Contains(dataFilter, "limit") {

		//}
		if strings.Contains(dataFilter, "pg=") {
			//	pgNums := strings.Split(dataFilter, "pg=")[1]
			//pgNum := strings.Split(pgNums, "&")[0]
			//dataPg = pgNum
		}
		limitData = " limit 20 OFFSET 1"
		dataFilterQ = whereQuery + dataFrom + dataTo + dataImei + dataPriority + dataEventId + dataSpeed +
			dataIgnition +
			dataGSMSignal +
			dataDigitalInput1 +
			dataDigitalOutput1 +
			dataSDStatus +
			dataDigitalInput2 +
			dataDigitalOutput2 +
			dataCrashDetection +
			dataOverSpeeding +
			dataExternalVoltage +
			dataBatteryVoltage +
			dataAnalogInput1 +
			dataAnalogInput2 +
			dataAnalogInput3 +
			dataAnalogInput4 +
			dataPCBTemperature +
			dataVehicleSpeed +
			dataEngineSpeed_RPM +
			dataEngineCoolantTemperature +
			dataFuellevelintank +
			dataCheckEngine +
			dataAirConditionPressureSwitch1 +
			dataAirConditionPressureSwitch2 +
			dataGearShiftindicator +
			dataDesiredGearValue +
			dataVehicleType +
			dataConditionimmobilizer +
			dataBrakePedalStatus +
			dataClutchPedalStatus +
			dataGearEngagedStatus +
			dataActualAccPedal +
			dataEngineThrottlePosition +
			dataIndicatedEngineTorque +
			dataEngineFrictionTorque +
			dataEngineActualTorque +
			dataCruiseControlOn_Off +
			dataSpeedLimiterOn_Off +
			dataConditioncruisecontrollamp +
			dataEngineFuleCutOff +
			dataConditioncatalystheatingactivated +
			dataACcompressorstatus +
			dataConditionmainrelay +
			datadistance +
			dataVirtualAccPedal +
			dataIntakeairtemperature +
			dataDesiredSpeed +
			dataOiltemperature_TCU +
			dataAmbientairtemperature +
			dataEMS_DTC +
			dataABS_DTC +
			dataBCM_DTC +
			dataACU_DTC +
			dataESC_DTC +
			dataICN_DTC +
			dataEPS_DTC +
			dataCAS_DTC +
			dataFCM_FN_DTC +
			dataICU_DTC +
			dataReserve_DTC +
			dataSensor1 +
			dataSensor2 +
			dataSensor3 +
			dataSensor4 +
			dataSensor5 +
			dataSensor6 +
			dataSensor7 +
			dataSensor8 +
			dataSensor9 +
			dataSensor10 +
			dataSensor11 +
			dataSensor12 +
			dataSensor13 +
			dataSensor14 +
			dataSensor15 +
			dataSensor16 +
			dataSensor17 +
			dataSensor18 +
			dataSensor19 +
			dataSensor20 +
			limitData
		// + " order by timestamp desc limit 1000 OFFSET " + dataPg
	}
	devicesLastPointsDatasQuery := "SELECT imei,timestamp AS ts,toUInt8(priority) AS priority,longitude,latitude,toInt32(altitude),toInt32(angle),toInt32(satellites),toInt32(speed), toUInt32(event_id),io_elements FROM avlpoints " + dataFilterQ

	rows, err := adb.GetChConn().Query(ctx, devicesLastPointsDatasQuery, &dataFilterQ)
	if err != nil {
		//logger.Info("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&1",
		//	zap.Any("2:", rows),
		//)
		return nil, err
	}
	logger.Info("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&2",
		zap.Any("3:", rows),
	)
	defer rows.Close()
	lastPointsData := make([]*devicepb.AVLData, 0)
	for rows.Next() {
		var (
			lastPointData = &devicepb.AVLData{}
			gps           = &devicepb.GPS{}
			priority      uint8
			elements      = make(map[string]float64)
		)
		err := rows.Scan(
			&lastPointData.Imei,
			&lastPointData.Timestamp,
			&priority,
			&gps.Longitude,
			&gps.Latitude,
			&gps.Altitude,
			&gps.Angle,
			&gps.Satellites,
			&gps.Speed,
			&lastPointData.EventId,
			&elements,
		)
		if err != nil {
			return nil, err
		}
		lastPointData.Gps = gps
		lastPointData.Priority = devicepb.PacketPriority(priority)
		for Name, Value := range elements {
			lastPointData.IoElements = append(lastPointData.IoElements, &devicepb.IOElement{
				ElementName:  Name,
				ElementValue: Value,
			})
		}
		lastPointsData = append(lastPointsData, lastPointData)
	}
	return lastPointsData, nil
}
