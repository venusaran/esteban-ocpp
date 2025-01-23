package chargepoint

import (
	"context"

	"gorm.io/gorm"

	"github.com/Beep-Technologies/beepbeep3-ocpp/api/rpc"
	"github.com/Beep-Technologies/beepbeep3-ocpp/internal/models"
	chargepoint "github.com/Beep-Technologies/beepbeep3-ocpp/internal/repository/charge_point"
	"github.com/Beep-Technologies/beepbeep3-ocpp/pkg/util"
)

type Service struct {
	db          *gorm.DB
	chargePoint chargepoint.BaseRepo
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db:          db,
		chargePoint: chargepoint.NewBaseRepo(db),
	}
}

func (srv Service) GetChargePoint(ctx context.Context, req *rpc.GetChargePointReq) (*rpc.GetChargePointResp, error) {
	cp, err := srv.chargePoint.GetChargePoint(ctx, req.EntityCode, req.ChargePointIdentifier)
	if err != nil {
		return nil, err
	}

	resChargePoint := &rpc.ChargePoint{}
	util.ConvertCopyStruct(resChargePoint, &cp, map[string]util.ConverterFunc{})

	res := &rpc.GetChargePointResp{
		ChargePoint: resChargePoint,
	}

	return res, nil
}

func (srv Service) GetChargePointIdTags(ctx context.Context, req *rpc.GetChargePointIdTagsReq) (*rpc.GetChargePointIdTagsResp, error) {
	cp, err := srv.chargePoint.GetChargePoint(ctx, req.EntityCode, req.ChargePointIdentifier)
	if err != nil {
		return nil, err
	}

	its, err := srv.chargePoint.GetIdTags(ctx, cp.ID)
	if err != nil {
		return nil, err
	}

	resIdTags := make([]*rpc.ChargePointIdTag, 0)
	for _, it := range its {
		resIdTag := rpc.ChargePointIdTag{}
		util.ConvertCopyStruct(&resIdTag, &it, map[string]util.ConverterFunc{})
		resIdTags = append(resIdTags, &resIdTag)
	}

	res := &rpc.GetChargePointIdTagsResp{
		ChargePointIdTags: resIdTags,
	}

	return res, nil
}

func (srv Service) UpdateChargePointDetails(ctx context.Context, req *rpc.UpdateChargePointDetailsReq) (*rpc.UpdateChargePointDetailsResp, error) {
	cp, err := srv.chargePoint.GetChargePoint(ctx, req.EntityCode, req.ChargePointIdentifier)
	if err != nil {
		return nil, err
	}

	cpID := cp.ID

	cp, err = srv.chargePoint.Update(
		ctx,
		cpID,
		[]string{
			"charge_point_vendor",
			"charge_point_model",
			"charge_point_serial_number",
			"charge_box_serial_number",
			"iccid",
			"imsi",
			"meter_type",
			"meter_serial_number",
			"firmware_version",
		},
		models.OcppChargePoint{
			ID:                      cpID,
			ChargePointVendor:       req.ChargePointVendor,
			ChargePointModel:        req.ChargePointModel,
			ChargePointSerialNumber: req.ChargePointSerialNumber,
			ChargeBoxSerialNumber:   req.ChargeBoxSerialNumber,
			Iccid:                   req.Iccid,
			Imsi:                    req.Imsi,
			MeterType:               req.MeterType,
			MeterSerialNumber:       req.MeterSerialNumber,
			FirmwareVersion:         req.FirmwareVersion,
		})

	if err != nil {
		return nil, err
	}

	resChargePoint := &rpc.ChargePoint{}
	util.ConvertCopyStruct(resChargePoint, cp, map[string]util.ConverterFunc{})
	res := &rpc.UpdateChargePointDetailsResp{
		ChargePoint: resChargePoint,
	}

	return res, nil
}

func (srv Service) GetChargePoints(ctx context.Context, req *rpc.GetChargePointsReq) (*rpc.GetChargePointsResp, error) {
	cps, err := srv.chargePoint.GetChargePoints(ctx, req.EntityCode)
	if err != nil {
		return nil, err
	}

	exCps := []*rpc.ChargePoint{}
	for _, cp := range cps {
		resChargePoint := &rpc.ChargePoint{
			Id:                      cp.ID,
			EntityCode:              cp.EntityCode,
			ChargePointIdentifier:   cp.ChargePointIdentifier,
			ChargePointVendor:       cp.ChargePointVendor,
			ChargePointModel:        cp.ChargePointModel,
			ChargePointSerialNumber: cp.ChargeBoxSerialNumber,
			ChargeBoxSerialNumber:   cp.ChargeBoxSerialNumber,
			Iccid:                   cp.Iccid,
			Imsi:                    cp.Imsi,
			MeterType:               cp.MeterType,
			MeterSerialNumber:       cp.MeterSerialNumber,
			FirmwareVersion:         cp.FirmwareVersion,
			ConnectorCount:          cp.ConnectorCount,
			OcppProtocol:            cp.OcppProtocol,
		}
		exCps = append(exCps, resChargePoint)
	}

	res := &rpc.GetChargePointsResp{
		ChargePoints: exCps,
	}

	return res, nil
}
